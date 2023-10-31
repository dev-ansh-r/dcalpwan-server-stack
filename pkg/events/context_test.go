package events_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

type testContextKeyType string

var testContextKey = testContextKeyType("ctx-test")

type testContextMarshaler struct{}

func (testContextMarshaler) MarshalContext(ctx context.Context) []byte {
	if val, ok := ctx.Value(testContextKey).(string); ok {
		return []byte(val)
	}
	return nil
}

func (testContextMarshaler) UnmarshalContext(ctx context.Context, b []byte) (context.Context, error) {
	return context.WithValue(ctx, testContextKey, string(b)), nil
}

func TestContextMarshaler(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	var m testContextMarshaler
	events.RegisterContextMarshaler("test", m)

	ctx := context.WithValue(context.Background(), testContextKey, "foo")

	evt := events.New(ctx, "test", "test")

	b, err := json.Marshal(evt)
	a.So(err, should.BeNil)

	unmarshaled, err := events.UnmarshalJSON(b)
	a.So(err, should.BeNil)
	a.So(unmarshaled, should.Resemble, evt)

	val, ok := unmarshaled.Context().Value(testContextKey).(string)
	if a.So(ok, should.BeTrue) {
		a.So(val, should.Equal, "foo")
	}
}
