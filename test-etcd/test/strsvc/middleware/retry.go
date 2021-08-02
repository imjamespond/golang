package middleware

import (
	"codechiev/utils"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"test-etcd/test/strsvc/model"
	"test-etcd/test/strsvc/service"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"

	// "golang.org/x/time/rate"
	httptransport "github.com/go-kit/kit/transport/http"
	jujuratelimit "github.com/juju/ratelimit"
)

func TestRetryMiddleware() model.Middleware {

	// Set some parameters for our client.
	var (
		qps         = 1                // beyond which we will return an error
		maxAttempts = 5                // per request, before giving up
		maxTime     = 10 * time.Second // wallclock time, before giving up
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
		e = MakeTestProxyEndpoint(_url)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e)   // 每个instance进行错误限流？
		e = ratelimit.NewDelayingLimiter(rate.NewLimiter(rate.Every(time.Second*5), qps))(e) // 每个instance进行延迟限流
		// e = NewTokenBucketLimitterWithJuju(jujuratelimit.NewBucket(time.Second, int64(qps)))(e)

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

var ErrLimitExceed = errors.New("rate limit exceed")

//NewTokenBucketLimitterWithJuju 使用juju/ratelimit创建限流中间件
func NewTokenBucketLimitterWithJuju(bkt *jujuratelimit.Bucket) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if bkt.TakeAvailable(1) == 0 {
				return nil, ErrLimitExceed
			}
			return next(ctx, request)
		}
	}
}

func MakeTestEndpoint(test service.Test) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		v, err := test(request.(model.Request))
		if err != nil {
			return model.Response{V: v, Err: err.Error()}, nil
		}
		return model.Response{V: "test", Err: ""}, nil
	}
}

func MakeTestProxyEndpoint(url *url.URL) endpoint.Endpoint {
	return httptransport.NewClient(
		"GET",
		url,
		func(_ context.Context, req *http.Request, _ interface{}) error {
			return nil
		},
		func(_ context.Context, resp *http.Response) (response interface{}, err error) {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			utils.ErrorIf(err)
			bodyString := string(bodyBytes)
			return bodyString, nil
		},
	).Endpoint()
}
