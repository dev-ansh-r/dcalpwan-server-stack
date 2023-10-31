package webmiddleware_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestMaxBody(t *testing.T) {
	m := MaxBody(16)

	t.Run("Normal Request", func(t *testing.T) {
		a := assertions.New(t)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusOK)
	})

	t.Run("Request Too Big", func(t *testing.T) {
		a := assertions.New(t)
		r := httptest.NewRequest(
			http.MethodPost, "/",
			bytes.NewBuffer([]byte("this is a little to much")),
		)
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := io.ReadAll(r.Body)
			a.So(err, should.HaveSameErrorDefinitionAs, ErrRequestBodyTooLarge)
			webhandlers.Error(w, r, err)
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusBadRequest)

		body, _ := io.ReadAll(res.Body)
		a.So(string(body), should.ContainSubstring, "request_body_too_large")
	})
}
