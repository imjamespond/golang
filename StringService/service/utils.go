package service

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

type GetRequestFunc func() (request interface{})

func DecodeRequest(getReq GetRequestFunc) httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (request interface{}, err error) {
		// var req interface{}
		req := getReq()
		// req := new(model.Request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			return nil, err
		}
		return req, nil
	}

}

func EncodeResponse() httptransport.EncodeResponseFunc {

	return func(_ context.Context, w http.ResponseWriter, response interface{}) error {
		return json.NewEncoder(w).Encode(response)
	}

}
