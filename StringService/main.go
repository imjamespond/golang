package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/imjamespond/string-service/controller"
	"github.com/imjamespond/string-service/model"
	"github.com/imjamespond/string-service/service"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.IStringService
	svc = service.StringService{}
	svc = service.LoggingMiddleware{logger, svc}

	uppercaseHandler := httptransport.NewServer(
		controller.MakeUppercaseEndpoint(svc),
		service.DecodeRequest(func() interface{} {
			return new(model.Request)
		}),
		service.EncodeResponse(),
	)

	countHandler := httptransport.NewServer(
		controller.MakeCountEndpoint(svc),
		service.DecodeRequest(func() interface{} {
			return new(model.Request)
		}),
		service.EncodeResponse(),
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	logger.Log(http.ListenAndServe(":8080", nil))
}
