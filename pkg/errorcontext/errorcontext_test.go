package errorcontext_test

import (
	"context"
	"errors"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errorcontext"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

var err error

func ExampleErrorContext() {
	ctx, cancel := errorcontext.New(test.Context())
	defer cancel(nil)

	go func() {
		for {
			// do work
			if err != nil {
				cancel(err)
			}
		}
	}()

	for {
		select {
		// case data := <-dataChan:
		case <-ctx.Done():
			return
		}
	}
}

func TestErrorContext(t *testing.T) {
	a := assertions.New(t)

	{
		err := errors.New("foo")
		ctx, cancel := errorcontext.New(test.Context())
		cancel(err)
		select {
		case <-ctx.Done():
			a.So(ctx.Err(), should.Equal, err)
		default:
			t.Error("Context was not done")
		}

		cancel(errors.New("other"))
		<-ctx.Done()
		a.So(ctx.Err(), should.Equal, err)
	}

	{
		ctx, cancel := context.WithCancel(test.Context())
		ctx, _ = errorcontext.New(ctx)
		cancel()
		select {
		case <-ctx.Done():
			a.So(ctx.Err(), should.Equal, context.Canceled)
		default:
			t.Error("Context was not done")
		}
	}
}
