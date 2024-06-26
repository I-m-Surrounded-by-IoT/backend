package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/host"
	"github.com/IBM/sarama"
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	kcircuitbreaker "github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	ggrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	ghttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/zijiren233/go-colorable"
	logkafka "github.com/zijiren233/logrus-kafka-hook"
	"github.com/zijiren233/stream"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type TCPHandler interface {
	ServeTcp(ctx context.Context, conn net.Conn) error
}

type TcpServer struct {
	l       net.Listener
	handler TCPHandler
	conf    *conf.TcpServer
}

func NewTcpServer(conf *conf.TcpServer, handler TCPHandler) *TcpServer {
	l, err := net.Listen("tcp", conf.Addr)
	if err != nil {
		logrus.Fatalf("failed to listen tcp: %v", err)
	}
	return &TcpServer{
		l:       l,
		handler: handler,
		conf:    conf,
	}
}

func (s *TcpServer) Start(ctx context.Context) error {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return err
		}
		select {
		case <-ctx.Done():
			return nil
		default:
			go func() {
				defer func() {
					if err := recover(); err != nil {
						logrus.Errorf("panic in tcp server: %v", err)
					}
				}()
				ctx, cancel := context.WithCancel(ctx)
				defer cancel()
				err := s.handler.ServeTcp(ctx, conn)
				if err != nil {
					if errors.Is(err, io.EOF) {
						return
					}
					logrus.Errorf("tcp server error: %v", err)
				}
			}()
		}
	}
}

func (s *TcpServer) Stop(ctx context.Context) error {
	return s.l.Close()
}

func (s *TcpServer) Endpoint() (*url.URL, error) {
	if s.conf.CustomEndpoint != "" {
		return url.Parse(s.conf.CustomEndpoint)
	}
	addr, err := host.Extract(s.conf.Addr, s.l)
	if err != nil {
		return nil, err
	}
	return url.Parse("tcp://" + addr)
}

func grpcHandlerFunc(gs *ggrpc.Server, other http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			gs.ServeHTTP(w, r)
		} else {
			other.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

type GrpcGatewayServer struct {
	gs   *ggrpc.Server
	hs   *ghttp.Server
	once sync.Once
}

func (s *GrpcGatewayServer) Start(ctx context.Context) error {
	s.once.Do(func() {
		s.hs.Handler = grpcHandlerFunc(s.gs, s.hs.Handler)
	})
	return s.hs.Start(ctx)
}

func (s *GrpcGatewayServer) Stop(ctx context.Context) error {
	return s.hs.Stop(ctx)
}

func (s *GrpcGatewayServer) Endpoint() (*url.URL, error) {
	return s.gs.Endpoint()
}

func (s *GrpcGatewayServer) GrpcRegistrar() grpc.ServiceRegistrar {
	return s.gs
}

func (s *GrpcGatewayServer) HttpRegistrar() *ghttp.Server {
	return s.hs
}

func (s *GrpcGatewayServer) Endpoints() ([]*url.URL, error) {
	ge, err := s.gs.Endpoint()
	if err != nil {
		return nil, err
	}
	he, err := s.hs.Endpoint()
	if err != nil {
		return nil, err
	}
	return []*url.URL{ge, he}, nil
}

func NewGrpcGatewayServer(config *conf.GrpcServerConfig) *GrpcGatewayServer {
	l, err := net.Listen("tcp", config.Addr)
	if err != nil {
		logrus.Fatalf("failed to listen tcp: %v", err)
	}

	middlewares := []middleware.Middleware{
		recovery.Recovery(),
		tracing.Server(),
	}

	if config.JwtSecret != "" {
		jwtSecret := []byte(config.JwtSecret)
		middlewares = append(middlewares, jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
			return jwtSecret, nil
		}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256)))
	}

	var hopts = []ghttp.ServerOption{
		ghttp.Middleware(middlewares...),
		ghttp.Listener(l),
		ghttp.Address(config.Addr),
	}

	var gopts = []ggrpc.ServerOption{
		ggrpc.Middleware(middlewares...),
		ggrpc.Listener(l),
		ggrpc.Address(config.Addr),
	}

	if config.Timeout != nil {
		hopts = append(hopts, ghttp.Timeout(config.Timeout.AsDuration()))
		gopts = append(gopts, ggrpc.Timeout(config.Timeout.AsDuration()))
	}

	var enableTls bool
	if config.Tls != nil && config.Tls.CertFile != "" && config.Tls.KeyFile != "" {
		enableTls = true
		var rootCAs *x509.CertPool
		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			logrus.Fatalf("failed to load system root CA: %v", err)
		}
		if config.Tls.CaFile != "" {
			b, err := os.ReadFile(config.Tls.CaFile)
			if err != nil {
				logrus.Fatalf("failed to read CA file: %v", err)
			}
			rootCAs.AppendCertsFromPEM(b)
		}

		cert, err := tls.LoadX509KeyPair(config.Tls.CertFile, config.Tls.KeyFile)
		if err != nil {
			logrus.Fatalf("failed to load cert and key: %v", err)
		}
		hopts = append(hopts, ghttp.TLSConfig(&tls.Config{
			RootCAs:      rootCAs,
			Certificates: []tls.Certificate{cert},
		}))
		gopts = append(gopts, ggrpc.TLSConfig(&tls.Config{
			RootCAs:      rootCAs,
			Certificates: []tls.Certificate{cert},
		}))
	}

	if config.CustomEndpoint != "" {
		u, err := url.Parse(config.CustomEndpoint)
		if err != nil {
			logrus.Fatalf("failed to parse custom endpoint: %v", err)
		}
		var (
			hu = *u
			gu = *u
		)
		if u.Scheme == "grpcs" || u.Scheme == "https" {
			hu.Scheme = "https"
			gu.Scheme = "grpcs"
		} else if u.Scheme == "grpc" || u.Scheme == "http" {
			hu.Scheme = "http"
			gu.Scheme = "grpc"
		} else if u.Scheme == "" {
			if enableTls {
				hu.Scheme = "https"
				gu.Scheme = "grpcs"
			} else {
				hu.Scheme = "http"
				gu.Scheme = "grpc"
			}
		} else {
			logrus.Fatalf("invalid custom endpoint scheme: %s", u.Scheme)
		}
		hopts = append(hopts, ghttp.Endpoint(&hu))
		gopts = append(gopts, ggrpc.Endpoint(&gu))
	}

	hs := ghttp.NewServer(hopts...)
	gs := ggrpc.NewServer(gopts...)
	return &GrpcGatewayServer{
		gs: gs,
		hs: hs,
	}
}

var (
	needColor     bool
	needColorOnce sync.Once
)

func ForceColor() bool {
	needColorOnce.Do(func() {
		if flags.DisableLogColor {
			needColor = false
			return
		}
		needColor = colorable.IsTerminal(os.Stdout.Fd())
	})
	return needColor
}

func GetEnvFiles(root string) ([]string, error) {
	var envs []string

	files, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), ".env") {
			envs = append(envs, file.Name())
		}
	}

	return envs, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GetRandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return stream.BytesToString(b)
}

type Backend struct {
	Endpoint  string
	Tls       bool
	JwtSecret string
	CustomCA  string
	TimeOut   string
}

type EtcdBackend struct {
	Backend
	ServiceName string
	Username    string
	Password    string
}

func NewEtcdGrpcConn(ctx context.Context, conf *EtcdBackend) (*grpc.ClientConn, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{conf.Endpoint},
		Username:  conf.Username,
		Password:  conf.Password,
	})
	if err != nil {
		return nil, err
	}
	if conf.ServiceName == "" {
		return nil, errors.New("new grpc client failed, service name is empty")
	}
	conf.Endpoint = fmt.Sprintf("discovery:///%s", conf.ServiceName)
	return NewDiscoveryGrpcConn(ctx, &conf.Backend, etcd.New(cli))
}

func NewDiscoveryGrpcConn(ctx context.Context, conf *Backend, d registry.Discovery) (*grpc.ClientConn, error) {
	if conf.Endpoint == "" {
		return nil, errors.New("new grpc client failed, endpoint is empty")
	}
	middlewares := []middleware.Middleware{
		kcircuitbreaker.Client(kcircuitbreaker.WithCircuitBreaker(func() circuitbreaker.CircuitBreaker {
			return sre.NewBreaker(
				sre.WithRequest(25),
				sre.WithWindow(time.Second*15),
			)
		})),
		tracing.Client(),
	}

	if conf.JwtSecret != "" {
		key := []byte(conf.JwtSecret)
		middlewares = append(middlewares, jwt.Client(func(token *jwtv5.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256)))
	}

	opts := []ggrpc.ClientOption{
		ggrpc.WithMiddleware(middlewares...),
		// ggrpc.WithOptions(grpc.WithBlock()),
	}

	if conf.TimeOut != "" {
		timeout, err := time.ParseDuration(conf.TimeOut)
		if err != nil {
			return nil, err
		}
		opts = append(opts, ggrpc.WithTimeout(timeout))
	}

	opts = append(opts, ggrpc.WithEndpoint(conf.Endpoint), ggrpc.WithDiscovery(d))

	var (
		con *grpc.ClientConn
		err error
	)
	if conf.Tls {
		var rootCAs *x509.CertPool
		rootCAs, err = x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		if conf.CustomCA != "" {
			rootCAs.AppendCertsFromPEM([]byte(conf.CustomCA))
		}
		opts = append(opts, ggrpc.WithTLSConfig(&tls.Config{
			RootCAs: rootCAs,
		}))

		con, err = ggrpc.Dial(
			ctx,
			opts...,
		)
	} else {
		con, err = ggrpc.DialInsecure(
			ctx,
			opts...,
		)
	}
	if err != nil {
		return nil, err
	}
	return con, nil
}

func NewSignalGrpcConn(ctx context.Context, conf *Backend) (*grpc.ClientConn, error) {
	if conf.Endpoint == "" {
		return nil, errors.New("new grpc client failed, endpoint is empty")
	}
	middlewares := []middleware.Middleware{
		kcircuitbreaker.Client(kcircuitbreaker.WithCircuitBreaker(func() circuitbreaker.CircuitBreaker {
			return sre.NewBreaker(
				sre.WithRequest(25),
				sre.WithWindow(time.Second*15),
			)
		})),
		tracing.Client(),
	}

	if conf.JwtSecret != "" {
		key := []byte(conf.JwtSecret)
		middlewares = append(middlewares, jwt.Client(func(token *jwtv5.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256)))
	}

	opts := []ggrpc.ClientOption{
		ggrpc.WithMiddleware(middlewares...),
		// ggrpc.WithOptions(grpc.WithBlock()),
	}

	if conf.TimeOut != "" {
		timeout, err := time.ParseDuration(conf.TimeOut)
		if err != nil {
			return nil, err
		}
		opts = append(opts, ggrpc.WithTimeout(timeout))
	}

	opts = append(opts, ggrpc.WithEndpoint(conf.Endpoint))

	var (
		con *grpc.ClientConn
		err error
	)
	if conf.Tls {
		var rootCAs *x509.CertPool
		rootCAs, err = x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		if conf.CustomCA != "" {
			rootCAs.AppendCertsFromPEM([]byte(conf.CustomCA))
		}
		opts = append(opts, ggrpc.WithTLSConfig(&tls.Config{
			RootCAs: rootCAs,
		}))

		con, err = ggrpc.Dial(
			ctx,
			opts...,
		)
	} else {
		con, err = ggrpc.DialInsecure(
			ctx,
			opts...,
		)
	}
	if err != nil {
		return nil, err
	}
	return con, nil
}

func SortUUID() string {
	return SortUUIDWithUUID(uuid.New())
}

func SortUUIDWithUUID(src uuid.UUID) string {
	dst := make([]byte, 32)
	hex.Encode(dst, src[:])
	return stream.BytesToString(dst)
}

type Logger struct {
	l *logrus.Logger
}

func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	var logrusLevel logrus.Level = logrus.InfoLevel
	switch level {
	case log.LevelDebug:
		logrusLevel = logrus.DebugLevel
	case log.LevelInfo:
		logrusLevel = logrus.InfoLevel
	case log.LevelWarn:
		logrusLevel = logrus.WarnLevel
	case log.LevelError:
		logrusLevel = logrus.ErrorLevel
	case log.LevelFatal:
		logrusLevel = logrus.FatalLevel
	}
	l.l.Log(logrusLevel, keyvals...)
	return nil
}

func TransLogrus(l *logrus.Logger) log.Logger {
	return &Logger{l: l}
}

func ForceNewKafkaClient(k *conf.KafkaConfig) sarama.Client {
	client, err := NewKafkaClient(k)
	if err != nil {
		log.Fatalf("failed to create kafka client: %v", err)
	}
	return client
}

func NewKafkaClient(k *conf.KafkaConfig) (sarama.Client, error) {
	if k == nil || k.Brokers == "" {
		return nil, errors.New("kafka config is empty")
	}
	opts := []logkafka.KafkaOptionFunc{
		logkafka.WithKafkaSASLHandshake(true),
		logkafka.WithKafkaSASLUser(k.User),
		logkafka.WithKafkaSASLPassword(k.Password),
	}
	if k.User != "" || k.Password != "" {
		opts = append(opts,
			logkafka.WithKafkaSASLEnable(true),
		)
	}
	client, err := logkafka.NewKafkaClient(
		strings.Split(k.Brokers, ","),
		opts...,
	)
	return client, err
}

func ValidateMac(mac string) (string, error) {
	if mac == "" {
		return "", fmt.Errorf("mac is empty")
	}
	if len(mac) == 17 {
		if mac[2] != ':' || mac[5] != ':' || mac[8] != ':' || mac[11] != ':' || mac[14] != ':' {
			return "", fmt.Errorf("mac is invalid")
		}
		return mac, nil
	} else if len(mac) == 12 {
		return fmt.Sprintf("%s:%s:%s:%s:%s:%s", mac[:2], mac[2:4], mac[4:6], mac[6:8], mac[8:10], mac[10:]), nil
	} else {
		return "", fmt.Errorf("mac is invalid")
	}
}

func DailKafka(k *conf.KafkaConfig, kafkaOpts ...logkafka.KafkaOptionFunc) (sarama.Client, error) {
	if k == nil || k.Brokers == "" {
		return nil, errors.New("kafka config is empty")
	}
	opts := []logkafka.KafkaOptionFunc{
		logkafka.WithKafkaSASLHandshake(true),
		logkafka.WithKafkaSASLUser(k.User),
		logkafka.WithKafkaSASLPassword(k.Password),
	}
	if k.User != "" || k.Password != "" {
		opts = append(opts,
			logkafka.WithKafkaSASLEnable(true),
		)
	}
	client, err := logkafka.NewKafkaClient(
		strings.Split(k.Brokers, ","),
		append(opts, kafkaOpts...)...,
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InitTracer(endpoint string, serverName string) error {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)))
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serverName),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
