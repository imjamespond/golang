package endpoint

import (
	"codechiev/utils"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"test-etcd/test/strsvc/model"
	"test-etcd/test/strsvc/service"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeTestEndpoint(test service.Test) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := test(request.(model.Request))
		if err != nil {
			return model.Response{V: v, Err: err.Error()}, nil
		}
		return model.Response{V: "test", Err: ""}, nil
	}
}

func MakeTestProxyEndpoint(url *url.URL) endpoint.Endpoint {
	return httptransport.NewClient(
		"GET",
		url,
		func(_ context.Context, req *http.Request, _ interface{}) error {
			return nil
		},
		func(_ context.Context, resp *http.Response) (response interface{}, err error) {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			utils.ErrorIf(err)
			bodyString := string(bodyBytes)
			return bodyString, nil
		},
	).Endpoint()
}
