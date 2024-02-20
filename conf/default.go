package conf

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

func DefaultGrpcServer() *GrpcServerConfig {
	return &GrpcServerConfig{
		Addr:    ":9000",
		Timeout: durationpb.New(time.Second * 15),
	}
}

func DefaultWebServer() *WebServerConfig {
	return &WebServerConfig{
		Addr: ":8080",
	}
}

func DefaultTcpServer() *TcpServer {
	return &TcpServer{
		Addr:    ":8000",
		Timeout: durationpb.New(time.Second * 15),
	}
}

func DefaultRegistry() *Registry {
	return &Registry{
		Etcd: &Registry_Etcd{},
	}
}

func DefaultKafka() *KafkaConfig {
	return &KafkaConfig{
		Brokers: "",
	}
}
