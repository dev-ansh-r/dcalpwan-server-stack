package test

import (
	"context"
)

// DefaultContext is the default context.
var DefaultContext = context.Background()

// Context returns DefaultContext.
func Context() context.Context {
	return DefaultContext
}
