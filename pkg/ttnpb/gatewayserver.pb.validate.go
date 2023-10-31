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

// ValidateFields checks the field values on GatewayUp with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GatewayUp) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GatewayUpFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "uplink_messages":

			for idx, item := range m.GetUplinkMessages() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return GatewayUpValidationError{
							field:  fmt.Sprintf("uplink_messages[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		case "gateway_status":

			if v, ok := interface{}(m.GetGatewayStatus()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GatewayUpValidationError{
						field:  "gateway_status",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "tx_acknowledgment":

			if v, ok := interface{}(m.GetTxAcknowledgment()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GatewayUpValidationError{
						field:  "tx_acknowledgment",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GatewayUpValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GatewayUpValidationError is the validation error returned by
// GatewayUp.ValidateFields if the designated constraints aren't met.
type GatewayUpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GatewayUpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GatewayUpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GatewayUpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GatewayUpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GatewayUpValidationError) ErrorName() string { return "GatewayUpValidationError" }

// Error satisfies the builtin error interface
func (e GatewayUpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGatewayUp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GatewayUpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GatewayUpValidationError{}

// ValidateFields checks the field values on GatewayDown with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GatewayDown) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = GatewayDownFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "downlink_message":

			if v, ok := interface{}(m.GetDownlinkMessage()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return GatewayDownValidationError{
						field:  "downlink_message",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return GatewayDownValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// GatewayDownValidationError is the validation error returned by
// GatewayDown.ValidateFields if the designated constraints aren't met.
type GatewayDownValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GatewayDownValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GatewayDownValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GatewayDownValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GatewayDownValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GatewayDownValidationError) ErrorName() string { return "GatewayDownValidationError" }

// Error satisfies the builtin error interface
func (e GatewayDownValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGatewayDown.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GatewayDownValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GatewayDownValidationError{}

// ValidateFields checks the field values on ScheduleDownlinkResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ScheduleDownlinkResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ScheduleDownlinkResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "delay":

			if m.GetDelay() == nil {
				return ScheduleDownlinkResponseValidationError{
					field:  "delay",
					reason: "value is required",
				}
			}

		case "downlink_path":

			if v, ok := interface{}(m.GetDownlinkPath()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return ScheduleDownlinkResponseValidationError{
						field:  "downlink_path",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		case "rx1":
			// no validation rules for Rx1
		case "rx2":
			// no validation rules for Rx2
		default:
			return ScheduleDownlinkResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ScheduleDownlinkResponseValidationError is the validation error returned by
// ScheduleDownlinkResponse.ValidateFields if the designated constraints
// aren't met.
type ScheduleDownlinkResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ScheduleDownlinkResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ScheduleDownlinkResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ScheduleDownlinkResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ScheduleDownlinkResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ScheduleDownlinkResponseValidationError) ErrorName() string {
	return "ScheduleDownlinkResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ScheduleDownlinkResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sScheduleDownlinkResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ScheduleDownlinkResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ScheduleDownlinkResponseValidationError{}

// ValidateFields checks the field values on ScheduleDownlinkErrorDetails with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ScheduleDownlinkErrorDetails) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = ScheduleDownlinkErrorDetailsFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "path_errors":

			for idx, item := range m.GetPathErrors() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return ScheduleDownlinkErrorDetailsValidationError{
							field:  fmt.Sprintf("path_errors[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return ScheduleDownlinkErrorDetailsValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// ScheduleDownlinkErrorDetailsValidationError is the validation error returned
// by ScheduleDownlinkErrorDetails.ValidateFields if the designated
// constraints aren't met.
type ScheduleDownlinkErrorDetailsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ScheduleDownlinkErrorDetailsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ScheduleDownlinkErrorDetailsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ScheduleDownlinkErrorDetailsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ScheduleDownlinkErrorDetailsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ScheduleDownlinkErrorDetailsValidationError) ErrorName() string {
	return "ScheduleDownlinkErrorDetailsValidationError"
}

// Error satisfies the builtin error interface
func (e ScheduleDownlinkErrorDetailsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sScheduleDownlinkErrorDetails.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ScheduleDownlinkErrorDetailsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ScheduleDownlinkErrorDetailsValidationError{}

// ValidateFields checks the field values on
// BatchGetGatewayConnectionStatsRequest with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *BatchGetGatewayConnectionStatsRequest) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = BatchGetGatewayConnectionStatsRequestFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "gateway_ids":

			if l := len(m.GetGatewayIds()); l < 1 || l > 100 {
				return BatchGetGatewayConnectionStatsRequestValidationError{
					field:  "gateway_ids",
					reason: "value must contain between 1 and 100 items, inclusive",
				}
			}

			for idx, item := range m.GetGatewayIds() {
				_, _ = idx, item

				if v, ok := interface{}(item).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return BatchGetGatewayConnectionStatsRequestValidationError{
							field:  fmt.Sprintf("gateway_ids[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		case "field_mask":

			if v, ok := interface{}(m.GetFieldMask()).(interface{ ValidateFields(...string) error }); ok {
				if err := v.ValidateFields(subs...); err != nil {
					return BatchGetGatewayConnectionStatsRequestValidationError{
						field:  "field_mask",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		default:
			return BatchGetGatewayConnectionStatsRequestValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// BatchGetGatewayConnectionStatsRequestValidationError is the validation error
// returned by BatchGetGatewayConnectionStatsRequest.ValidateFields if the
// designated constraints aren't met.
type BatchGetGatewayConnectionStatsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchGetGatewayConnectionStatsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchGetGatewayConnectionStatsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchGetGatewayConnectionStatsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchGetGatewayConnectionStatsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchGetGatewayConnectionStatsRequestValidationError) ErrorName() string {
	return "BatchGetGatewayConnectionStatsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BatchGetGatewayConnectionStatsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchGetGatewayConnectionStatsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchGetGatewayConnectionStatsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchGetGatewayConnectionStatsRequestValidationError{}

// ValidateFields checks the field values on
// BatchGetGatewayConnectionStatsResponse with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *BatchGetGatewayConnectionStatsResponse) ValidateFields(paths ...string) error {
	if m == nil {
		return nil
	}

	if len(paths) == 0 {
		paths = BatchGetGatewayConnectionStatsResponseFieldPathsNested
	}

	for name, subs := range _processPaths(append(paths[:0:0], paths...)) {
		_ = subs
		switch name {
		case "entries":

			for key, val := range m.GetEntries() {
				_ = val

				// no validation rules for Entries[key]

				if v, ok := interface{}(val).(interface{ ValidateFields(...string) error }); ok {
					if err := v.ValidateFields(subs...); err != nil {
						return BatchGetGatewayConnectionStatsResponseValidationError{
							field:  fmt.Sprintf("entries[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						}
					}
				}

			}

		default:
			return BatchGetGatewayConnectionStatsResponseValidationError{
				field:  name,
				reason: "invalid field path",
			}
		}
	}
	return nil
}

// BatchGetGatewayConnectionStatsResponseValidationError is the validation
// error returned by BatchGetGatewayConnectionStatsResponse.ValidateFields if
// the designated constraints aren't met.
type BatchGetGatewayConnectionStatsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BatchGetGatewayConnectionStatsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BatchGetGatewayConnectionStatsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BatchGetGatewayConnectionStatsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BatchGetGatewayConnectionStatsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BatchGetGatewayConnectionStatsResponseValidationError) ErrorName() string {
	return "BatchGetGatewayConnectionStatsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e BatchGetGatewayConnectionStatsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBatchGetGatewayConnectionStatsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BatchGetGatewayConnectionStatsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BatchGetGatewayConnectionStatsResponseValidationError{}
