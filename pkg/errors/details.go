package errors

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

type detailer interface {
	Details() []proto.Message
}

func (e *Error) addDetails(details ...proto.Message) {
	e.details = append(e.details, details...)
	if e.stack == nil {
		e.stack = callers(4)
	}
	e.clearGRPCStatus()
}

// WithDetails returns the error with the given details set.
// This appends to any existing details in the Error.
func (e *Error) WithDetails(details ...proto.Message) *Error {
	if e == nil {
		return e
	}
	dup := *e
	dup.addDetails(details...)
	return &dup
}

// WithDetails returns a new error from the definition, and sets the given details.
func (d *Definition) WithDetails(details ...proto.Message) *Error {
	if d == nil {
		return nil
	}
	e := build(d, 0) // Don't refactor this to build(...).WithDetails(...)
	e.addDetails(details...)
	return e
}

// Details of the error. Usually structs from ttnpb or google.golang.org/genproto/googleapis/rpc/errdetails.
func (e *Error) Details() (details []proto.Message) {
	if e == nil {
		return nil
	}
	if e.cause != nil {
		details = append(details, Details(e.cause)...)
	}
	if len(e.details) > 0 {
		details = append(details, e.details...)
	}
	return
}

// Details are not present in the error definition, so this just returns nil.
func (*Definition) Details() []proto.Message { return nil }

// Details gets the details of the error.
func Details(err error) []proto.Message {
	if detailsErr := (detailer)(nil); errors.As(err, &detailsErr) {
		return detailsErr.Details()
	}
	if ttnErr, ok := From(err); ok {
		return ttnErr.Details()
	}
	return nil
}
