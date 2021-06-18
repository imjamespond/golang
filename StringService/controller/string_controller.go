package controller

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-kit/kit/endpoint"

	"github.com/imjamespond/string-service/model"
	"github.com/imjamespond/string-service/service"
)

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func MakeUppercaseEndpoint(svc service.IStringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*model.Request)
		fmt.Println(req.Body)
		v, err := svc.Uppercase(req.Body)
		if err != nil {
			return model.Response{v, err.Error()}, nil
		}
		return model.Response{v, ""}, nil
	}
}

func MakeCountEndpoint(svc service.IStringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(*model.Request)
		fmt.Println(req.Body)
		v := svc.Count(req.Body)
		return model.Response{strconv.Itoa(v), ""}, nil
	}
}
