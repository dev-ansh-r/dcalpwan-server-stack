package errors_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestResembles(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	a.So(errors.Resemble(nil, nil), should.BeTrue)

	defInvalidArgument := errors.DefineInvalidArgument("test_resembles_invalid_argument", "invalid argument")

	// Nil errors never resemble.
	a.So(errors.Resemble(defInvalidArgument, nil), should.BeFalse)
	a.So(errors.Resemble(nil, defInvalidArgument), should.BeFalse)

	// Typed nil is invalid.
	a.So(errors.Resemble(defInvalidArgument, (*errors.Definition)(nil)), should.BeFalse)
	a.So(errors.Resemble(defInvalidArgument, (*errors.Error)(nil)), should.BeFalse)

	errInvalidArgument := defInvalidArgument.WithAttributes("foo", "bar")
	grpcErrInvalidArgument := errInvalidArgument.GRPCStatus().Err()

	// Errors and definitions resemble:
	a.So(errors.Resemble(errInvalidArgument, defInvalidArgument), should.BeTrue)

	// Should resemble gRPC error:
	a.So(errors.Resemble(grpcErrInvalidArgument, defInvalidArgument), should.BeTrue)
	a.So(errors.Resemble(grpcErrInvalidArgument, errInvalidArgument), should.BeTrue)

	defPermissionDenied := errors.DefinePermissionDenied("test_resembles_permission_denied", "permission denied")

	a.So(errors.Resemble(defInvalidArgument, defPermissionDenied), should.BeFalse)

	defWrapper := errors.Define("wrapper", "something went wrong")

	a.So(errors.Resemble(
		defWrapper.WithCause(defPermissionDenied),
		defWrapper,
	), should.BeTrue)
	a.So(errors.Resemble(
		defWrapper.WithCause(defPermissionDenied),
		defWrapper.WithCause(defInvalidArgument),
	), should.BeTrue)
}
