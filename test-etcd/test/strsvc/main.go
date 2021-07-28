package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"test-etcd/test/strsvc/endpoint"
	"test-etcd/test/strsvc/middleware"
	"test-etcd/test/strsvc/service"
	"test-etcd/test/strsvc/util"
)

// Transports expose the service to the network. In this first example we utilize JSON over HTTP.
func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.IStringService
	svc = service.StringService{}
	svc = middleware.LoggingMiddleware{Logger: logger, Next: svc}

	uppercaseHandler := httptransport.NewServer(
		endpoint.MakeUppercaseEndpoint(svc),
		util.DecodeRequest(),
		util.EncodeResponse(),
	)

	countHandler := httptransport.NewServer(
		endpoint.MakeCountEndpoint(svc),
		util.DecodeRequest(),
		util.EncodeResponse(),
	)

	var _svc service.Service
	_svc = service.Foobar{}
	_svc = middleware.LoggingSvcMiddleware{Logger: logger, Next: _svc}
	foobarHandler := httptransport.NewServer(
		endpoint.MakeEndpoint(_svc),
		util.DecodeRequest(),
		util.EncodeResponse(),
	)

	mw := middleware.TestLoggingMiddleware(log.With(logger, "method", "test1"))
	ep := mw(endpoint.MakeTestEndpoint(service.CallTest))
	mw = middleware.TestLoggingMiddleware(log.With(logger, "method", "test2"))
	ep = mw(ep)
	testHandler := httptransport.NewServer(
		ep,
		util.DecodeRequest(),
		util.EncodeResponse(),
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.Handle("/foobar", foobarHandler)
	http.Handle("/test", testHandler)
	logger.Log(http.ListenAndServe(":8080", nil))
}
