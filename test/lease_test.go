package test

import (
	"context"
	"log"
	"testing"
	"utils"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestLease(t *testing.T) {
	cli, err := GetClient()
	utils.FatalIf(err)
	defer cli.Close()

	// 创建一个5秒的租约
	resp, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatal(err)
	}

	// 5秒钟之后, 这个key就会被移除
	_, err = cli.Put(context.TODO(), "foo", "bar", clientv3.WithLease(resp.ID))
	if err != nil {
		log.Fatal(err)
	}
}
