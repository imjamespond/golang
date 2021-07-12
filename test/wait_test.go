package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWg(t *testing.T) {
	const threadNum = 4
	wg := sync.WaitGroup{}
	wg.Add(threadNum)
	for i := 1; i <= threadNum; i++ {
		go func(i int) {
			fmt.Printf("Thread %d begin\n", i)
			time.Sleep(time.Second * time.Duration(i))
			fmt.Printf("Thread %d done\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("All Threads finished\n")
}
