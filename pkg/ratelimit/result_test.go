package ratelimit_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/ratelimit"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestResult(t *testing.T) {
	r := ratelimit.Result{Limit: 10, Remaining: 5, RetryAfter: time.Second, ResetAfter: 2 * time.Second}

	t.Run("Zero", func(t *testing.T) {
		a := assertions.New(t)
		a.So(r.IsZero(), should.BeFalse)
		a.So(ratelimit.Result{}.IsZero(), should.BeTrue)
	})

	t.Run("Headers", func(t *testing.T) {
		t.Run("HTTP", func(t *testing.T) {
			a := assertions.New(t)
			headers := http.Header{}
			r.SetHTTPHeaders(headers)
			a.So(headers.Get("X-Rate-Limit-Limit"), should.Equal, "10")
			a.So(headers.Get("X-Rate-Limit-Available"), should.Equal, "5")
			a.So(headers.Get("X-Rate-Limit-Retry"), should.Equal, "1")
			a.So(headers.Get("X-Rate-Limit-Reset"), should.Equal, "2")
		})

		t.Run("gRPC", func(t *testing.T) {
			a := assertions.New(t)
			headers := r.GRPCHeaders()

			a.So(headers.Get("X-Rate-Limit-Limit"), should.Resemble, []string{"10"})
			a.So(headers.Get("X-Rate-Limit-Available"), should.Resemble, []string{"5"})
			a.So(headers.Get("X-Rate-Limit-Retry"), should.Resemble, []string{"1"})
			a.So(headers.Get("X-Rate-Limit-Reset"), should.Resemble, []string{"2"})
		})
	})
}
