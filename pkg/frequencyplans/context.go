package frequencyplans

import "context"

type fallbackKeyType struct{}

var fallbackKey fallbackKeyType

// WithFallbackID returns a derived context with the given frequency plan ID to be used as fallback.
func WithFallbackID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, fallbackKey, id)
}

// FallbackIDFromContext returns the fallback frequency plan ID and whether it's set using WithFallbackID.
func FallbackIDFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(fallbackKey).(string)
	return id, ok
}
