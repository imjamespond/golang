package utils

import (
	"log"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := After(1, func() {
		log.Println("...")
	})
	time.Sleep(5500 * time.Millisecond)
	timer.Stop()
	time.Sleep(5 * time.Second)
}

// time.After 例子
func After(seconds int64, do func()) *time.Ticker {
	ticker := time.NewTicker(time.Second * time.Duration(seconds)) // 启动定时器

	go func(t *time.Ticker) {
		for range t.C {
			do()
		}
	}(ticker)

	// 不再使用了，结束它
	// ticker.Stop()
	return ticker
}
