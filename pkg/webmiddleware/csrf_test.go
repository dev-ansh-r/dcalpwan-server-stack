package webmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/csrf"
	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/auth"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestCSRF(t *testing.T) {
	a := assertions.New(t)
	authKey := []byte("1234123412341234123412341234123412341234123412341234123412341234")
	m := CSRF(authKey)

	t.Run("Protects non-idempotent methods when using a Session Token", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/", nil)
		r.Header.Set("Authorization", "Bearer "+auth.JoinToken(auth.SessionToken, "XXX", "YYY"))
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusForbidden)
	})

	t.Run("Allows non-idempotent methods when using an API Key", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPost, "/", nil)
		r.Header.Set("Authorization", "Bearer "+auth.JoinToken(auth.APIKey, "XXX", "YYY"))
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusOK)
	})

	t.Run("Allows access with valid CSRF token", func(t *testing.T) {
		var csrfToken string
		var r *http.Request

		// Obtain CSRF token.
		r = httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			csrfToken = csrf.Token(r)
		})).ServeHTTP(rec, r)
		res := rec.Result()

		cookies := res.Cookies()
		a.So(cookies, should.HaveLength, 1)

		// Make request
		r = httptest.NewRequest(http.MethodPost, "/", nil)
		r.Header.Set("X-CSRF-Token", csrfToken)
		r.AddCookie(cookies[0])
		rec = httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})).ServeHTTP(rec, r)
		res = rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusOK)
	})
}
