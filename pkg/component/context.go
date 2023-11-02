package component

import (
	"context"
	"time"
)

type crossContext struct {
	cancelCtx context.Context
	valueCtx  context.Context
}

// Deadline implements context.Context using the cancel context.
func (ctx *crossContext) Deadline() (deadline time.Time, ok bool) {
	return ctx.cancelCtx.Deadline()
}

// Done implements context.Context using the cancel context.
func (ctx *crossContext) Done() <-chan struct{} {
	return ctx.cancelCtx.Done()
}

// Err implements context.Context using the cancel context.
func (ctx *crossContext) Err() error {
	return ctx.cancelCtx.Err()
}

// Value implements context.Context using the value context.
func (ctx *crossContext) Value(key any) any {
	return ctx.valueCtx.Value(key)
}
