package test

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
	"utils"

	concurrency "go.etcd.io/etcd/client/v3/concurrency"
)

func TestLock(t *testing.T) {

	count := 0
	wg := sync.WaitGroup{}
	wg.Add(4)

	// 创建两个单独的会话用来演示锁竞争
	test := func(id int) {
		cli, err := GetClient()
		utils.FatalIf(err)
		defer cli.Close() // filo

		sess, err := concurrency.NewSession(cli)
		utils.FatalIf(err)
		defer sess.Close()

		mutex := concurrency.NewMutex(sess, "/foo-lock")
		for {
			err = mutex.Lock(context.TODO())
			utils.FatalIf(err)

			if count++; count <= 100 {
				log.Println("lock for ", id, count)
			} else {
				break
			}

			err = mutex.Unlock(context.TODO())
			utils.FatalIf(err)

			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}

		wg.Done()
	}
	go test(1)
	go test(2)
	go test(3)
	go test(4)

	wg.Wait()
}
