package ratelimit_test

import "go.thethings.network/lorawan-stack/v3/pkg/ratelimit"

var (
	maxRate      uint = 10
	overrideRate uint = 15
)

type mockLimiter struct {
	calledWithResource ratelimit.Resource

	limit  bool
	result ratelimit.Result
}

func (l *mockLimiter) RateLimit(resource ratelimit.Resource) (bool, ratelimit.Result) {
	l.calledWithResource = resource
	return l.limit, l.result
}

type muxMockLimiter map[string]*mockLimiter

func (l muxMockLimiter) RateLimit(resource ratelimit.Resource) (bool, ratelimit.Result) {
	for _, class := range resource.Classes() {
		if limiter, ok := l[class]; ok {
			return limiter.RateLimit(resource)
		}
	}
	return false, ratelimit.Result{}
}
