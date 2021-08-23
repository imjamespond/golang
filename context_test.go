package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestCtxTimeout(t *testing.T) {
	// 创建一个超时时间为100毫秒的上下文
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// 创建一个访问 nc -l 127.0.0.1 8000
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8000", nil)
	// 将超时上下文关联到创建的请求上
	req = req.WithContext(ctx)

	// 创建一个HTTP客户端并执行请求
	client := &http.Client{}
	res, err := client.Do(req)
	// 如果请求失败了，记录到STDOUT
	if err != nil {
		fmt.Println("Request failed:", err) // context deadline exceeded
		return
	}
	// 请求成功后打印状态码
	fmt.Println("Response received, status code:", res.StatusCode)
}

func TestCtxCancel(t *testing.T) {
	operation1 := func(ctx context.Context) error {
		time.Sleep(100 * time.Millisecond)
		return errors.New("failed")
	}

	operation2 := func(ctx context.Context) {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("done")
		case <-ctx.Done():
			fmt.Println("halted operation2")
		}
	}
	// 新建一个上下文
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		operation2(ctx)
	}()

	err := operation1(ctx)
	if err != nil {
		cancel() // cancel operation2 if there is an error return by oper1
	}
}
