package testetcd

import (
	"context"
	"fmt"
	"testing"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func Test_(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.1.107:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	cli.Put(ctx, "greeting", "Hello world")
	resp, err := cli.Get(ctx, "greeting")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp.Kvs)
	defer cli.Close()
	defer cancel()
}
