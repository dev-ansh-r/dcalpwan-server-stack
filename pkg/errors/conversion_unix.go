//go:build linux || darwin
// +build linux darwin

package errors

import (
	"errors"
	"syscall"
)

func syscallErrorAttributes(err error) []any {
	if matched := (syscall.Errno)(0); errors.As(err, &matched) {
		// syscall.Errono do not contain any sensitive information and are safe to render.
		return []any{
			"error", matched.Error(),
			"timeout", matched.Timeout(),
		}
	}
	return nil
}
