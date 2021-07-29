package middleware

import (
	"codechiev/utils"
	"net/url"
	makeEndpoint "test-etcd/test/strsvc/endpoint"
	"test-etcd/test/strsvc/model"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

func TestRetryMiddleware() model.Middleware {

	// Set some parameters for our client.
	var (
		qps         = 1                      // beyond which we will return an error
		maxAttempts = 3                      // per request, before giving up
		maxTime     = 250 * time.Millisecond // wallclock time, before giving up
	)

	// Otherwise, construct an endpoint for each instance in the list, and add
	// it to a fixed set of endpoints. In a real service, rather than doing this
	// by hand, you'd probably use package sd's support for your service
	// discovery system.
	var (
		instanceList = []string{"https://vv.video.qq.com/checktime", "https://api.muxiaoguo.cn/api/tianqi?city=长沙&type=1"}
		endpointer   sd.FixedEndpointer
	)

	for _, instance := range instanceList {
		_url, err := url.Parse(instance)
		utils.PanicIf(err)
		var e endpoint.Endpoint
		e = makeEndpoint.MakeTestProxyEndpoint(_url)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(3*time.Second), qps))(e)
		endpointer = append(endpointer, e)
	}

	// Now, build a single, retrying, load-balancing endpoint out of all of
	// those individual endpoints.
	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(maxAttempts, maxTime, balancer)

	// And finally, return the Middleware
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		// Endpoint
		return retry
	}
}
