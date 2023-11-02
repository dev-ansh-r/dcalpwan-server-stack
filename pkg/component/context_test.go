package component_test

import (
	"context"
	"testing"
	"time"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestContextDecoupling(t *testing.T) {
	a := assertions.New(t)

	c, err := component.New(test.GetLogger(t), &component.Config{})
	a.So(err, should.BeNil)

	dl := time.Now().Add(1 * time.Second)

	ctx := test.Context()
	ctx = context.WithValue(ctx, "key", "value")
	ctx, cancel := context.WithDeadline(ctx, dl)
	defer cancel()

	ctx = c.FromRequestContext(ctx)
	deadline, set := ctx.Deadline()
	a.So(deadline, should.NotEqual, dl)
	a.So(set, should.BeFalse)

	select {
	case <-ctx.Done():
		t.Fatal("Context expired")
	case <-time.After(time.Second):
		a.So(ctx.Err(), should.BeNil)
	}

	val := ctx.Value("key")
	strVal, ok := val.(string)
	a.So(ok, should.BeTrue)
	a.So(strVal, should.Equal, "value")
}
