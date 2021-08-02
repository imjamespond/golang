package main

import (
	"context"
	"net/http"
	"os"
	"test-etcd/test/echo"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var (
		//etcd服務地址
		etcdServer = "test:2377"
		//服務的資訊目錄
		prefix = "/services/echo/"
		//當前啟動服務範例的地址
		instance = "http://127.0.0.1:50052"
		//服務範例註冊的路徑
		key = prefix + instance
		//服務範例註冊的val
		value = instance
		ctx   = context.Background()
		//服務監聽地址
		serviceAddress = ":50052"
	)
	//etcd的連線引數
	options := etcdv3.ClientOptions{
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 3,
	}
	//建立etcd連線
	client, err := etcdv3.NewClient(ctx, []string{etcdServer}, options)
	if err != nil {
		panic(err)
	}
	// 建立註冊器
	registrar := etcdv3.NewRegistrar(client, etcdv3.Service{
		Key:   key,
		Value: value,
	}, log.NewNopLogger())
	// 註冊器啟動註冊
	registrar.Register()

	testHandler := httptransport.NewServer(
		// Endpointer
		func(_ context.Context, request interface{}) (interface{}, error) {
			return request, nil
		},
		echo.DecodeRequestFunc,
		echo.EncodeResponseFunc,
	)

	logger := log.NewLogfmtLogger(os.Stderr)
	http.Handle(prefix, testHandler)
	logger.Log(http.ListenAndServe(serviceAddress, nil))
}
