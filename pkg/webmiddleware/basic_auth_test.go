package webmiddleware_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestBasicAuth(t *testing.T) {
	m := BasicAuth("Password Protected", AuthUser("username", "password"))

	t.Run("No Auth", func(t *testing.T) {
		a := assertions.New(t)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("Handler was called when it shouldn't have")
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusUnauthorized)
		a.So(res.Header.Get("WWW-Authenticate"), should.ContainSubstring, "Password Protected")
	})

	t.Run("No Username or Password", func(t *testing.T) {
		a := assertions.New(t)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.SetBasicAuth("", "")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("Handler was called when it shouldn't have")
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusUnauthorized)
		a.So(res.Header.Get("WWW-Authenticate"), should.ContainSubstring, "Password Protected")
	})

	t.Run("No Username", func(t *testing.T) {
		a := assertions.New(t)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.SetBasicAuth("", "password")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("Handler was called when it shouldn't have")
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusUnauthorized)
		a.So(res.Header.Get("WWW-Authenticate"), should.ContainSubstring, "Password Protected")
	})

	t.Run("No Password", func(t *testing.T) {
		a := assertions.New(t)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.SetBasicAuth("username", "")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("Handler was called when it shouldn't have")
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusUnauthorized)
		a.So(res.Header.Get("WWW-Authenticate"), should.ContainSubstring, "Password Protected")
	})

	t.Run("Wrong Auth", func(t *testing.T) {
		a := assertions.New(t)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.SetBasicAuth("username", "wrong-password")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t.Error("Handler was called when it shouldn't have")
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusUnauthorized)
		a.So(res.Header.Get("WWW-Authenticate"), should.ContainSubstring, "Password Protected")
	})

	t.Run("Successful Auth", func(t *testing.T) {
		a := assertions.New(t)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		r.SetBasicAuth("username", "password")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Secret"))
		})).ServeHTTP(rec, r)
		res := rec.Result()
		a.So(res.StatusCode, should.Equal, http.StatusOK)
		body, _ := io.ReadAll(res.Body)
		a.So(string(body), should.Equal, "Secret")
	})
}
