package test

import (
	"context"
	"log"
	"testing"
	"utils"

	"test-etcd/common"
)

func Test1(t *testing.T) {
	cli, err := GetClient()
	utils.FatalIf(err)
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	resp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	log.Println(resp)

	common.HandleErr(err)
}
