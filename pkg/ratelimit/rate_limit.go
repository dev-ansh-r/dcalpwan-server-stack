package ratelimit

import (
	"context"

	"github.com/throttled/throttled"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
)

// Interface can be used to rate limit access to a Resource.
type Interface interface {
	// RateLimit limits access on a Resource.
	//
	// The RateLimit operations returns true if the rate limit for the requested resource has been exceeded, along
	// with metadata for the rate limiting operation.
	RateLimit(resource Resource) (limit bool, result Result)
}

// NoopRateLimiter does not enforce any rate limits.
type NoopRateLimiter struct{}

// RateLimit implements ratelimit.Interface.
func (*NoopRateLimiter) RateLimit(Resource) (bool, Result) {
	return false, Result{}
}

type rateLimiter struct {
	ctx     context.Context
	limiter throttled.RateLimiter
}

// RateLimit implements ratelimit.Interface.
func (l *rateLimiter) RateLimit(resource Resource) (bool, Result) {
	ok, result, err := l.limiter.RateLimit(resource.Key(), 1)
	if err != nil {
		// NOTE: The memstore.MemStore implementation does not fail.
		log.FromContext(l.ctx).Error("Rate limiter failed")
		return true, Result{}
	}

	return ok, Result{
		Limit:      result.Limit,
		Remaining:  result.Remaining,
		RetryAfter: result.RetryAfter,
		ResetAfter: result.ResetAfter,
	}
}

// muxRateLimiter is a ratelimit.Interface that supports multiple rate limiting profiles.
// If no rate limiting profile is set for a resource class, then no rate limits are applied.
type muxRateLimiter struct {
	defaultLimiter Interface
	limiters       map[string]Interface
}

// RateLimit implements ratelimit.Interface.
func (l *muxRateLimiter) RateLimit(resource Resource) (bool, Result) {
	for _, c := range resource.Classes() {
		if limiter, ok := l.limiters[c]; ok {
			return limiter.RateLimit(resource)
		}
	}
	return l.defaultLimiter.RateLimit(resource)
}

// Require checks that the rate limit for a Resource has not been exceeded.
func Require(limiter Interface, resource Resource) error {
	if limit, result := limiter.RateLimit(resource); limit {
		return errRateLimitExceeded.WithAttributes(
			"key", resource.Key(),
			"rate", result.Limit,
		)
	}
	return nil
}

// RateLimitKeyer can be implemented by request messages. If implemented, the
// string returned by the RateLimitKey() method is appended to the key used for
// rate limiting.
type RateLimitKeyer interface {
	RateLimitKey() string
}
