package main

import (
	"net/http"
	"os"

	gkEndpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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

	var ep gkEndpoint.Endpoint
	mw := middleware.TestLoggingMiddleware(log.With(logger, "method", "test1"))
	// ep := mw(endpoint.MakeTestEndpoint(service.CallTest))
	ep = mw(middleware.TestRetryMiddleware()(ep))
	mw = middleware.TestLoggingMiddleware(log.With(logger, "method", "test2"))
	ep = mw(ep)
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	mw = middleware.InstrumentingMiddleware(requestCount, requestLatency)
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
	http.Handle("/metrics", promhttp.Handler())
	logger.Log(http.ListenAndServe(":8080", nil))
}
