package endpoint

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-kit/kit/endpoint"

	"test-etcd/test/strsvc/model"
	"test-etcd/test/strsvc/service"
)

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our middleware interface)

// curl -d '{"foo":"bar"}' localhost:8080/uppercase
func MakeUppercaseEndpoint(svc service.IStringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Request)
		fmt.Println(req)
		foo := req["foo"]
		if nil == foo {
			return model.Response{V: "", Err: "foo is nil"}, nil
		}
		v, err := svc.Uppercase(foo.(string))
		if err != nil {
			return model.Response{V: v, Err: err.Error()}, nil
		}
		return model.Response{V: v, Err: ""}, nil
	}
}

// curl -d '{"hello":"AlinaLopez"}' localhost:8080/count
func MakeCountEndpoint(svc service.IStringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.Request)
		fmt.Println(req)
		v := svc.Count(req["hello"].(string))
		return model.Response{V: strconv.Itoa(v), Err: ""}, nil
	}
}

func MakeEndpoint(svc service.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.Do(request.(model.Request))
		if err != nil {
			return model.Response{V: v, Err: err.Error()}, nil
		}
		return model.Response{V: v, Err: ""}, nil
	}
}
