package component

import "go.thethings.network/lorawan-stack/v3/pkg/ratelimit"

func (c *Component) RateLimiter() ratelimit.Interface {
	return c.limiter
}
