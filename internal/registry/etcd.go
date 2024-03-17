package registry

import (
	"strings"

	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func newEtcd(c *conf.Registry_Etcd) *clientv3.Client {
	if c == nil || c.Endpoint == "" {
		return nil
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(c.Endpoint, ","),
		Username:    c.Username,
		Password:    c.Password,
		DialTimeout: c.Timeout.AsDuration(),
	})
	if err != nil {
		logrus.Fatalf("failed to create etcd client: %v", err)
	}
	return client
}
