package main

import (
	"bytes"
	"codechiev/utils"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"

	httptransport "github.com/go-kit/kit/transport/http"
)

var (
	//註冊中心地址
	etcdServer = "test:2377"
	//監聽的服務字首
	prefix = "/services/echo/"
	ctx    = context.Background()
)

func main() {

	options := etcdv3.ClientOptions{
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 3,
	}
	//連線註冊中心
	client, err := etcdv3.NewClient(ctx, []string{etcdServer}, options)
	if err != nil {
		panic(err)
	}
	logger := log.NewNopLogger()
	//建立範例管理器, 此管理器會Watch監聽etc中prefix的目錄變化更新快取的服務範例資料
	instancer, err := etcdv3.NewInstancer(client, prefix, logger)
	if err != nil {
		panic(err)
	}
	//建立端點管理器， 此管理器根據Factory和監聽的到範例建立endPoint並訂閱instancer的變化動態更新Factory建立的endPoint
	endpointer := sd.NewEndpointer(instancer, reqFactory, logger)
	//建立負載均衡器
	balancer := lb.NewRoundRobin(endpointer)
	/**
	  我們可以通過負載均衡器直接獲取請求的endPoint，發起請求
	  reqEndPoint,_ := balancer.Endpoint()
	*/
	/**
	  也可以通過retry定義嘗試次數進行請求
	*/
	reqEndPoint := lb.Retry(3, 30*time.Second, balancer)
	//現在我們可以通過 endPoint 發起請求了
	req := []byte("foobar") // struct{}{}
	_ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	reqEndPoint(_ctx, req)
	time.Sleep(30 * time.Second)
	cancel()

	// curl -d '{"hello":"AlinaLopez"}' localhost:8080/test
	// testHandler := httptransport.NewServer(
	// 	func(ctx context.Context, request interface{}) (interface{}, error) {
	// 		return reqEndPoint(ctx, request)
	// 	},
	// 	echo.DecodeRequestFunc,
	// 	echo.EncodeResponseFunc,
	// )
	// http.Handle("/test", testHandler)
	// logger.Log(http.ListenAndServe(":8080", nil))
}

//通過傳入的 範例地址  建立對應的請求endPoint
func reqFactory(instanceAddr string) (endpoint.Endpoint, io.Closer, error) {
	fmt.Println("請求服務: ", instanceAddr)
	_url, _ := url.Parse(instanceAddr + prefix)
	return httptransport.NewClient(
		"GET", _url,
		func(_ context.Context, req *http.Request, request interface{}) error {
			req.Body = ioutil.NopCloser(bytes.NewBuffer(request.([]byte)))
			return nil
		},
		func(_ context.Context, resp *http.Response) (response interface{}, err error) {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			utils.ErrorIf(err)
			bodyString := string(bodyBytes)
			return bodyString, nil
		},
	).Endpoint(), nil, nil
}
