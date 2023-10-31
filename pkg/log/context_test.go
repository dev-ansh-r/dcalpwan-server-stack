package log_test

import (
	"context"
	"testing"

	"github.com/smarty/assertions"
	. "go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestContext(t *testing.T) {
	a := assertions.New(t)
	ctx := context.Background()

	logger := NewLogger(NoopHandler)

	a.So(FromContext(ctx), should.NotBeNil)
	a.So(FromContext(ctx), should.Equal, Noop)

	ctx = NewContext(ctx, logger)

	a.So(FromContext(ctx), should.Equal, logger)

	t.Run("NewContextWithField", func(t *testing.T) {
		a := assertions.New(t)
		withKV := FromContext(NewContextWithField(ctx, "key", "value")).(Entry)
		fields := withKV.Fields().Fields()
		v, ok := fields["key"]
		a.So(ok, should.BeTrue)
		a.So(v, should.Equal, "value")
	})

	t.Run("NewContextWithFields", func(t *testing.T) {
		a := assertions.New(t)
		withKV := FromContext(NewContextWithFields(ctx, Fields("key", "value"))).(Entry)
		fields := withKV.Fields().Fields()
		v, ok := fields["key"]
		a.So(ok, should.BeTrue)
		a.So(v, should.Equal, "value")
	})
}
