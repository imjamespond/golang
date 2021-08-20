package test

import (
	"fmt"
	"testing"
	"utils"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestTransaction(t *testing.T) {
	cli, err := GetClient()
	utils.FatalIf(err)
	defer cli.Close() // filo

	txnrsp, err := cli.Txn(cli.Ctx()).If(
		clientv3.Compare(clientv3.Value("/mydir/test1"), "=", "test1"),
	).Then(
		clientv3.OpPut("/mydir/test1", "text1"),
		clientv3.OpPut("/mydir/test2", "text2"),
	).Else(
		clientv3.OpPut("/mydir/test1", "test1"),
		clientv3.OpPut("/mydir/test2", "test2"),
	).Commit()
	utils.FatalIf(err)

	for _, rp := range txnrsp.Responses {
		fmt.Println(rp.GetResponse()) // revision
	}

	txnrsp, err = cli.Txn(cli.Ctx()).If().Then(
		clientv3.OpGet("/mydir/test1"),
		clientv3.OpGet("/mydir/test2"),
	).Commit()
	utils.FatalIf(err)

	for _, rp := range txnrsp.Responses {
		for _, ev := range rp.GetResponseRange().Kvs {
			fmt.Println(ev)
		}
	}

}
