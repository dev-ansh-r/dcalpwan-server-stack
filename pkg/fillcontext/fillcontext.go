// Package fillcontext implements a mechanism to fill contexts.
package fillcontext

import "context"

// Filler extends the context
type Filler func(context.Context) context.Context
