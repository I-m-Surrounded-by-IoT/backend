package utils

import (
	"context"
	"crypto/tls"
	"crypto/x509"
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
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-kratos/aegis/circuitbreaker/sre"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	kcircuitbreaker "github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	ggrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	ghttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/sirupsen/logrus"
	"github.com/zijiren233/go-colorable"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	jwtv4 "github.com/golang-jwt/jwt/v4"
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
		panic(err)
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

func NewGrpcGatewayServer(config *conf.GrpcServer) *GrpcGatewayServer {
	middlewares := []middleware.Middleware{recovery.Recovery()}
	if config.JwtSecret != "" {
		jwtSecret := []byte(config.JwtSecret)
		middlewares = append(middlewares, jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
			return jwtSecret, nil
		}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)))
	}

	l, err := net.Listen("tcp", config.Addr)
	if err != nil {
		panic(err)
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
			fmt.Println("Failed to load system root CA:", err)
			panic(err)
		}
		if config.Tls.CaFile != "" {
			b, err := os.ReadFile(config.Tls.CaFile)
			if err != nil {
				panic(err)
			}
			rootCAs.AppendCertsFromPEM(b)
		}

		cert, err := tls.LoadX509KeyPair(config.Tls.CertFile, config.Tls.KeyFile)
		if err != nil {
			panic(err)
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
			panic(err)
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
			panic("invalid custom endpoint scheme: " + u.Scheme)
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

func GetRand[T any](s []T) (v T) {
	if len(s) == 0 {
		return
	}
	return s[rand.Intn(len(s))]
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
	middlewares := []middleware.Middleware{kcircuitbreaker.Client(kcircuitbreaker.WithCircuitBreaker(func() circuitbreaker.CircuitBreaker {
		return sre.NewBreaker(
			sre.WithRequest(25),
			sre.WithWindow(time.Second*15),
		)
	}))}

	if conf.JwtSecret != "" {
		key := []byte(conf.JwtSecret)
		middlewares = append(middlewares, jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)))
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
	middlewares := []middleware.Middleware{kcircuitbreaker.Client(kcircuitbreaker.WithCircuitBreaker(func() circuitbreaker.CircuitBreaker {
		return sre.NewBreaker(
			sre.WithRequest(25),
			sre.WithWindow(time.Second*15),
		)
	}))}

	if conf.JwtSecret != "" {
		key := []byte(conf.JwtSecret)
		middlewares = append(middlewares, jwt.Client(func(token *jwtv4.Token) (interface{}, error) {
			return key, nil
		}, jwt.WithSigningMethod(jwtv4.SigningMethodHS256)))
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
