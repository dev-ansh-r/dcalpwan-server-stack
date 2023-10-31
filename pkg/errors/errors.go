// Package errors implements rich errors that carry contextual information
// such as stack traces, causality and attributes.
package errors

import (
	"encoding/hex"
	"fmt"
	"sync/atomic"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"
)

// New returns an error that formats as the given text.
// This way of creating errors should be avoided if possible.
func New(text string) *Error {
	return build(&Definition{
		namespace:     namespace(2),
		messageFormat: text,
		code:          uint32(codes.Unknown),
	}, 4)
}

// Error is a rich error implementation.
type Error struct {
	*Definition
	*stack
	correlationID string
	cause         error
	details       []proto.Message
	attributes    map[string]any
	grpcStatus    *atomic.Value
}

func (e *Error) String() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("error:%s (%s)", e.FullName(), e.FormatMessage(e.PublicAttributes()))
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return e.String()
}

// Fields implements the log.Fielder interface.
func (e *Error) Fields() map[string]any {
	if e == nil {
		return nil
	}
	res := make(map[string]any)
	pref := "error_cause"
	for cause := Cause(e); cause != nil; cause = Cause(cause) {
		res[pref] = cause.Error()
		pref += "_cause"
	}
	for k, v := range Attributes(Stack(e)...) {
		res[k] = v
	}
	return res
}

// Interface is the interface of an error.
type Interface interface {
	DefinitionInterface
	Attributes() map[string]any
	Cause() error
	Details() (details []proto.Message)
}

var generateCorrelationIDs = true

// GenerateCorrelationIDs configures whether random correlation IDs are generated
// for each error. This is enabled by default.
func GenerateCorrelationIDs(enable bool) {
	generateCorrelationIDs = enable
}

// build an error from the definition, skipping the first frames of the call stack.
func build(d *Definition, skip int) *Error {
	e := Error{
		Definition: d,
		grpcStatus: new(atomic.Value),
	}
	if generateCorrelationIDs {
		e.correlationID = hex.EncodeToString(uuid.NewV4().Bytes())
	}
	if skip > 0 {
		e.stack = callers(skip)
	}
	return &e
}
