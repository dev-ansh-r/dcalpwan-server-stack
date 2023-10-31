package scripting

import (
	"context"
)

// Engine represents a scripting engine.
type Engine interface {
	Run(ctx context.Context, script, fn string, params ...any) (func(target any) error, error)
}

// AheadOfTimeEngine extends Engine with the capability of compiling the script ahead of time.
type AheadOfTimeEngine interface {
	Engine

	Compile(ctx context.Context, script string) (run func(context.Context, string, ...any) (func(any) error, error), err error)
}
