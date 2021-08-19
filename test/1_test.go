package test

import (
	"context"
	"log"
	"testing"
	"utils"

	"test-etcd/common"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func Test1(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{Hosts},
		DialTimeout: dialTimeout,
	})

	utils.PanicIf(err)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	log.Println(resp)

	common.HandleErr(err)

	defer cli.Close()
}
