package errors

import (
	"runtime"

	"github.com/pkg/errors"
)

type stack []uintptr

// callers returns the call stack, skipping the first frames.
func callers(skip int) *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

// StackTrace implements the errors.stackTracer interface.
func (s *stack) StackTrace() errors.StackTrace {
	if s == nil {
		return nil
	}
	f := make([]errors.Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = errors.Frame((*s)[i])
	}
	return f
}
