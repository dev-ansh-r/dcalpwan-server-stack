package api_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/packages/loradms/v1/api"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func chanRoundTrip(reqChan chan<- *http.Request, respChan <-chan *http.Response, errChan <-chan error) http.RoundTripper {
	return roundTripperFunc(
		func(req *http.Request) (*http.Response, error) {
			reqChan <- req
			return <-respChan, <-errChan
		},
	)
}

func withClient(ctx context.Context, t *testing.T, opts []api.Option, f func(context.Context, *testing.T, <-chan *http.Request, chan<- *http.Response, chan<- error, *api.Client)) {
	reqChan := make(chan *http.Request, 5)
	respChan := make(chan *http.Response, 5)
	errChan := make(chan error, 5)
	cl, err := api.New(&http.Client{
		Transport: chanRoundTrip(reqChan, respChan, errChan),
	}, opts...)
	if !assertions.New(t).So(err, should.BeNil) {
		t.FailNow()
	}
	f(ctx, t, reqChan, respChan, errChan, cl)
}

func TestNoAuth(t *testing.T) {
	withClient(test.Context(), t, nil,
		func(ctx context.Context, t *testing.T, reqChan <-chan *http.Request, respChan chan<- *http.Response, errChan chan<- error, cl *api.Client) {
			a := assertions.New(t)

			respChan <- &http.Response{
				Body: io.NopCloser(bytes.NewBufferString("")),
			}
			errChan <- nil

			resp, err := cl.Do(ctx, "foo", "bar", "baz", http.MethodGet, nil)
			req := <-reqChan
			a.So(resp, should.NotBeNil)
			a.So(err, should.BeNil)
			a.So(req.Header, should.NotContainKey, "Authorization")
		})
}

func TestAuth(t *testing.T) {
	withClient(test.Context(), t, []api.Option{api.WithToken("foobar")},
		func(ctx context.Context, t *testing.T, reqChan <-chan *http.Request, respChan chan<- *http.Response, errChan chan<- error, cl *api.Client) {
			a := assertions.New(t)

			respChan <- &http.Response{
				Body: io.NopCloser(bytes.NewBufferString("")),
			}
			errChan <- nil

			resp, err := cl.Do(ctx, "foo", "bar", "baz", http.MethodGet, nil)
			req := <-reqChan
			a.So(resp, should.NotBeNil)
			a.So(err, should.BeNil)
			a.So(req.Header, should.ContainKey, "Authorization")
			a.So(req.Header["Authorization"], should.Resemble, []string{"foobar"})
		})
}
