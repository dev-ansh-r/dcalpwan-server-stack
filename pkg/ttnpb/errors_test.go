package ttnpb_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestGRPCConversion(t *testing.T) {
	a := assertions.New(t)

	errDef := errors.Define("test_grpc_conversion_err_def", "gRPC Conversion Error", "foo")
	a.So(errors.FromGRPCStatus(errDef.GRPCStatus()).Definition, should.EqualErrorOrDefinition, errDef)

	errHello := errDef.WithAttributes("foo", "bar", "baz", "qux")
	errHelloExpected := errDef.WithAttributes("foo", "bar")
	a.So(errors.FromGRPCStatus(errHello.GRPCStatus()), should.EqualErrorOrDefinition, errHelloExpected)
}
