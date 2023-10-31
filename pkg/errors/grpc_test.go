package errors_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	_ "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestGRPCConversion(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	errCause := errors.Define("test_grpc_conversion_err_cause", "gRPC Conversion Error Cause")

	errDef := errors.Define("test_grpc_conversion_err_def", "gRPC Conversion Error", "foo")
	a.So(errors.FromGRPCStatus(errDef.GRPCStatus()).Definition, should.EqualErrorOrDefinition, errDef)

	errHello := errDef.WithAttributes("foo", "bar", "baz", "qux").WithCause(errCause)
	errHelloExpected := errDef.WithAttributes("foo", "bar").WithCause(errCause)
	a.So(errors.FromGRPCStatus(errHello.GRPCStatus()), should.EqualErrorOrDefinition, errHelloExpected)
}
