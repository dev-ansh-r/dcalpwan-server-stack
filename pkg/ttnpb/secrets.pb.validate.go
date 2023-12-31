// Code generated by protoc-gen-fieldmask. DO NOT EDIT.

package ttnpb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// ValidateFields checks the field values on Secret with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Secret) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = SecretFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "key_id":
			// no validation rules for KeyId
		case "value":

			if len(m.GetValue()) > 2048 {
				return SecretValidationError{
					field:  "value",
					reason: "value length must be at most 2048 bytes",
				}
			}

		default:
			return SecretValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// SecretValidationError is the validation error returned by
// Secret.ValidateFields if the designated constraints aren't met.
type SecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecretValidationError) ErrorName() string { return "SecretValidationError" }

// Error satisfies the builtin error interface
func (e SecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecretValidationError{}
