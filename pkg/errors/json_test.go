package errors_test

import (
	"encoding/json"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	_ "go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestJSONConversion(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	errDef := errors.Define("test_json_conversion_err_def", "JSON Conversion Error", "foo")

	b, err := json.Marshal(errDef)
	a.So(err, should.BeNil)

	var unmarshaledDef errors.Definition
	err = json.Unmarshal(b, &unmarshaledDef)
	a.So(err, should.BeNil)
	a.So(&unmarshaledDef, should.EqualErrorOrDefinition, errDef)

	errHello := errDef.WithAttributes("foo", "bar", "baz", "qux")
	errHelloExpected := errDef.WithAttributes("foo", "bar")

	b, err = json.Marshal(errHello)
	a.So(err, should.BeNil)

	var unmarshaled errors.Error
	err = json.Unmarshal(b, &unmarshaled)
	a.So(err, should.BeNil)
	a.So(&unmarshaled, should.EqualErrorOrDefinition, errHelloExpected)
}
