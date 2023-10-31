package errors

import (
	"errors"
)

// Alias standard library error functions.
var (
	As     = errors.As
	Is     = errors.Is
	Unwrap = errors.Unwrap
)

// Unwrap makes the Error implement error unwrapping.
func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

// Is makes the Error implement error comparison.
func (e *Error) Is(target error) bool {
	if e == nil {
		return target == nil
	}
	if namedErr := (interface {
		Namespace() string
		Name() string
	})(nil); errors.As(target, &namedErr) {
		return namedErr.Namespace() == e.Namespace() &&
			namedErr.Name() == e.Name()
	}
	return false
}

// Unwrap makes the Definition implement error unwrapping.
func (*Definition) Unwrap() error {
	return nil
}

// Is makes the Definition implement error comparison.
func (d *Definition) Is(target error) bool {
	if d == nil {
		return target == nil
	}
	if namedErr := (interface {
		Namespace() string
		Name() string
	})(nil); errors.As(target, &namedErr) {
		return namedErr.Namespace() == d.Namespace() &&
			namedErr.Name() == d.Name()
	}
	return false
}
