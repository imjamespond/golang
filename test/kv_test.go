package test

import (
	"context"
	"fmt"
	"test-etcd/common"
	"testing"
	"utils"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestKV(t *testing.T) {
	cli, err := GetClient()
	utils.FatalIf(err)
	defer cli.Close() // filo

	_, err = cli.Put(context.Background(), "/mydir/foo", "bar")
	common.HandleErr(err)
	_, err = cli.Put(context.Background(), "/mydir/hello", "world")
	common.HandleErr(err)

	resp, err := cli.KV.Get(context.Background(), "/mydir/", clientv3.WithPrefix()) // greater than clientv3.WithMinCreateRev(5916)
	common.HandleErr(err)
	for i, kv := range resp.Kvs {
		fmt.Println(i, string(kv.Key), "-", string(kv.Value), "/////////", kv)
	}

}
