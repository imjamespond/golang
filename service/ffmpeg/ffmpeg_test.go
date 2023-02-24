package ffmpeg

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	svc := GetInstance()
	svc.Start()
	for i := 0; i < 8; i++ {
		go func(num int) {
			for i := 0; i < 1; i++ {
				str := fmt.Sprintf("%d-%d", num, i)
				svc.Add(&Job{Type: 0, Data: str})
				fmt.Println("add", str)
			}
		}(i)
	}
	time.Sleep(time.Second * 10)
}

func TestInstance(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go func(i int) {
			fmt.Println("go", i)
			GetInstance()
			fmt.Println("end", i)
			wg.Done()
		}(i)
	}
	fmt.Println("wait")
	wg.Wait()
	fmt.Println("finished")
}
