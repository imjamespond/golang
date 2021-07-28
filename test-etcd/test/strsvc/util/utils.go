package util

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"test-etcd/test/strsvc/model"

	httptransport "github.com/go-kit/kit/transport/http"
)

func DecodeRequest() httptransport.DecodeRequestFunc {
	return func(_ context.Context, r *http.Request) (request interface{}, err error) {
		var req model.Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty string")
