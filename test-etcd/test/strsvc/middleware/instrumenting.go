package middleware

import (
	"context"
	"test-etcd/test/strsvc/model"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics"
)

func InstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram) model.Middleware {
	// Middleware
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		// Endpoint
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			defer func(begin time.Time) {
				lvs := []string{"method", "some method", "error", "some error"}
				requestCount.With(lvs...).Add(1)
				requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return next(ctx, request)
		}
	}
}
