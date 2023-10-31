// Package scripting provides a generic abstraction layer for running scripts at runtime.
package scripting

import "time"

// Options contains engine options.
type Options struct {
	StackDepthLimit int
	Timeout         time.Duration
}

// DefaultOptions are the default Options.
var DefaultOptions = Options{
	StackDepthLimit: 32,
	Timeout:         100 * time.Millisecond,
}
