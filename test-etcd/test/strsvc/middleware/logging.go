package middleware

import (
	"context"
	"test-etcd/test/strsvc/model"
	"test-etcd/test/strsvc/service"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   service.IStringService
}

func (mw LoggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Uppercase(s) //调用service
	return
}

func (mw LoggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.Next.Count(s)
	return
}

type LoggingSvcMiddleware struct {
	Logger log.Logger
	Next   service.Service
}

func (mw LoggingSvcMiddleware) Do(req model.Request) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Do(req) //调用service
	return
}

func TestLoggingMiddleware(logger log.Logger) model.Middleware {
	// Middleware
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		// Endpoint
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
