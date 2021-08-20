package test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"utils"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var wg sync.WaitGroup

func TestWatcher(t *testing.T) {
	wg.Add(2)

	watchingKey := "/mydir/"
	// 模拟KV的变化
	// go func() {
	// 	for {
	// 		_, err = cli.Put(context.TODO(), watchingKey, "helios1")
	// 		_, err = cli.Delete(context.TODO(), watchingKey)
	// 		time.Sleep(2 * time.Second)
	// 	}
	// }()

	watch := func(num int) {
		cli, err := GetClient()
		utils.FatalIf(err)
		defer cli.Close()

		rch := cli.Watch(context.Background(), watchingKey, clientv3.WithPrefix())
		fmt.Println("blocked...")
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("%d: %s %q : %q, %d, %d\n", num, ev.Type, ev.Kv.Key, ev.Kv.Value, ev.Kv.ModRevision, ev.Kv.Version)
			}
		}
		wg.Done()
	}
	go watch(1)
	go watch(2)

	wg.Wait()
}
