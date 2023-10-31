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

// ValidateFields checks the field values on ALCSyncCommand with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ALCSyncCommand) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ALCSyncCommandFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "cid":

			if _, ok := ALCSyncCommandIdentifier_name[int32(m.GetCid())]; !ok {
				return ALCSyncCommandValidationError{
					field:  "cid",
					reason: "value must be one of the defined enum values",
				}
			}

		case "payload":
			if len(subs) == 0 {
				subs = []string{
					"app_time_req", "app_time_ans",
				}
			}
			for name, subs := range _processPaths(subs) {
				_ = subs
				switch name {
				case "app_time_req":
					w, ok := m.Payload.(*ALCSyncCommand_AppTimeReq_)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetAppTimeReq()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return ALCSyncCommandValidationError{
								field:  "app_time_req",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				case "app_time_ans":
					w, ok := m.Payload.(*ALCSyncCommand_AppTimeAns_)
					if !ok || w == nil {
						continue
					}

					if v, ok := interface{}(m.GetAppTimeAns()).(interface{ ValidateFields(...string) error }); ok {
						if err := v.ValidateFields(subs...); err != nil {
							return ALCSyncCommandValidationError{
								field:  "app_time_ans",
								reason: "embedded message failed validation",
								cause:  err,
							}
						}
					}

				}
			}
		default:
			return ALCSyncCommandValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ALCSyncCommandValidationError is the validation error returned by
// ALCSyncCommand.ValidateFields if the designated constraints aren't met.
type ALCSyncCommandValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ALCSyncCommandValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ALCSyncCommandValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ALCSyncCommandValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ALCSyncCommandValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ALCSyncCommandValidationError) ErrorName() string { return "ALCSyncCommandValidationError" }

// Error satisfies the builtin error interface
func (e ALCSyncCommandValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sALCSyncCommand.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ALCSyncCommandValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ALCSyncCommandValidationError{}

// ValidateFields checks the field values on ALCSyncCommand_AppTimeReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ALCSyncCommand_AppTimeReq) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ALCSyncCommand_AppTimeReqFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "DeviceTime":

			if m.GetDeviceTime() == nil {
				return ALCSyncCommand_AppTimeReqValidationError{
					field:  "DeviceTime",
					reason: "value is required",
				}
			}

		case "TokenReq":

			if m.GetTokenReq() > 255 {
				return ALCSyncCommand_AppTimeReqValidationError{
					field:  "TokenReq",
					reason: "value must be less than or equal to 255",
				}
			}

		case "AnsRequired":
			// no validation rules for AnsRequired
		default:
			return ALCSyncCommand_AppTimeReqValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ALCSyncCommand_AppTimeReqValidationError is the validation error returned by
// ALCSyncCommand_AppTimeReq.ValidateFields if the designated constraints
// aren't met.
type ALCSyncCommand_AppTimeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ALCSyncCommand_AppTimeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ALCSyncCommand_AppTimeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ALCSyncCommand_AppTimeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ALCSyncCommand_AppTimeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ALCSyncCommand_AppTimeReqValidationError) ErrorName() string {
	return "ALCSyncCommand_AppTimeReqValidationError"
}

// Error satisfies the builtin error interface
func (e ALCSyncCommand_AppTimeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sALCSyncCommand_AppTimeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ALCSyncCommand_AppTimeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ALCSyncCommand_AppTimeReqValidationError{}

// ValidateFields checks the field values on ALCSyncCommand_AppTimeAns with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ALCSyncCommand_AppTimeAns) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ALCSyncCommand_AppTimeAnsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "TimeCorrection":
			// no validation rules for TimeCorrection
		case "TokenAns":

			if m.GetTokenAns() > 255 {
				return ALCSyncCommand_AppTimeAnsValidationError{
					field:  "TokenAns",
					reason: "value must be less than or equal to 255",
				}
			}

		default:
			return ALCSyncCommand_AppTimeAnsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ALCSyncCommand_AppTimeAnsValidationError is the validation error returned by
// ALCSyncCommand_AppTimeAns.ValidateFields if the designated constraints
// aren't met.
type ALCSyncCommand_AppTimeAnsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ALCSyncCommand_AppTimeAnsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ALCSyncCommand_AppTimeAnsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ALCSyncCommand_AppTimeAnsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ALCSyncCommand_AppTimeAnsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ALCSyncCommand_AppTimeAnsValidationError) ErrorName() string {
	return "ALCSyncCommand_AppTimeAnsValidationError"
}

// Error satisfies the builtin error interface
func (e ALCSyncCommand_AppTimeAnsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sALCSyncCommand_AppTimeAns.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ALCSyncCommand_AppTimeAnsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ALCSyncCommand_AppTimeAnsValidationError{}
