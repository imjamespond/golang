package test

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	Host1          = "test:2381"
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

func GetClient() (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{Host1},
		DialTimeout: dialTimeout,
	})
}
