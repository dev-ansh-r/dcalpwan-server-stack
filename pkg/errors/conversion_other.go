//go:build !linux && !darwin
// +build !linux,!darwin

package errors

func syscallErrorAttributes(error) []any {
	return nil
}
