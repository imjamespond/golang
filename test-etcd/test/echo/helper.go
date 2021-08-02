package echo

import (
	"codechiev/utils"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DecodeRequestFunc(_ context.Context, req *http.Request) (request interface{}, err error) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	utils.ErrorIf(err)
	fmt.Println("req:", string(bodyBytes))
	return bodyBytes, nil
}
func EncodeResponseFunc(_ context.Context, w http.ResponseWriter, response interface{}) error {
	switch o := response.(type) {
	case []byte:
		w.Write(response.([]byte))
	case string:
		w.Write([]byte(response.(string)))
	default:
		fmt.Println("unknown type of response", o)
	}
	return nil
}
