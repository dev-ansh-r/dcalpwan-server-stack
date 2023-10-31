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

func TestCORS(t *testing.T) {
	a := assertions.New(t)

	config := CORSConfig{
		AllowedHeaders:   []string{"X-Allowed-Header"},
		AllowedMethods:   []string{http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedOrigins:   []string{"http://localhost"},
		ExposedHeaders:   []string{"X-Exposed-Header"},
		MaxAge:           600,
		AllowCredentials: true,
	}

	m := CORS(config)

	t.Run("Cross Origin Request", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPut, "/", nil)
		r.Header.Set("Origin", "http://localhost")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("CORS-Enabled"))
		})).ServeHTTP(rec, r)
		res := rec.Result()

		a.So(res.StatusCode, should.Equal, http.StatusOK)
		a.So(res.Header.Get("Access-Control-Allow-Origin"), should.Equal, "http://localhost")
		a.So(res.Header.Get("Access-Control-Expose-Headers"), should.ContainSubstring, "X-Exposed-Header")
		a.So(res.Header.Get("Access-Control-Allow-Credentials"), should.Equal, "true")

		body, _ := io.ReadAll(res.Body)
		a.So(string(body), should.Equal, "CORS-Enabled")
	})

	t.Run("Preflight OK", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodOptions, "/", nil)
		r.Header.Set("Origin", "http://localhost")
		r.Header.Set("Access-Control-Request-Method", http.MethodPut)
		r.Header.Set("Access-Control-Request-Headers", "X-Allowed-Header")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("CORS-Enabled"))
		})).ServeHTTP(rec, r)
		res := rec.Result()

		a.So(res.StatusCode, should.BeBetweenOrEqual, 200, 299)
		a.So(res.Header.Get("Access-Control-Allow-Origin"), should.Equal, "http://localhost")
		a.So(res.Header.Get("Access-Control-Allow-Methods"), should.ContainSubstring, http.MethodPut)
		a.So(res.Header.Get("Access-Control-Allow-Headers"), should.ContainSubstring, "X-Allowed-Header")
		a.So(res.Header.Get("Access-Control-Max-Age"), should.Equal, "600")
		a.So(res.Header.Get("Access-Control-Allow-Credentials"), should.Equal, "true")

		body, _ := io.ReadAll(res.Body)
		a.So(body, should.BeEmpty)
	})

	t.Run("Preflight Fail Origin", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodOptions, "/", nil)
		r.Header.Set("Origin", "http://notlocalhost")
		r.Header.Set("Access-Control-Request-Method", http.MethodPut)
		r.Header.Set("Access-Control-Request-Headers", "X-Allowed-Header")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("CORS-Enabled"))
		})).ServeHTTP(rec, r)
		res := rec.Result()

		a.So(res.StatusCode, should.BeBetweenOrEqual, 200, 299)
		a.So(res.Header.Get("Access-Control-Allow-Origin"), should.BeEmpty)

		body, _ := io.ReadAll(res.Body)
		a.So(body, should.BeEmpty)
	})

	t.Run("Preflight Fail Method", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodOptions, "/", nil)
		r.Header.Set("Origin", "http://localhost")
		r.Header.Set("Access-Control-Request-Method", http.MethodPatch)
		r.Header.Set("Access-Control-Request-Headers", "X-Allowed-Header")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("CORS-Enabled"))
		})).ServeHTTP(rec, r)
		res := rec.Result()

		// NOTE: Not sure if it is correct behavior to return 405 here.
		// The preflight request is successful, so you'd expect a 2xx code.
		// As long as the headers are okay, I think it's fine.
		a.So(res.StatusCode, should.Equal, http.StatusMethodNotAllowed)
		a.So(res.Header.Get("Access-Control-Allow-Origin"), should.BeEmpty)

		body, _ := io.ReadAll(res.Body)
		a.So(body, should.BeEmpty)
	})

	t.Run("Preflight Fail Header", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodOptions, "/", nil)
		r.Header.Set("Origin", "http://localhost")
		r.Header.Set("Access-Control-Request-Method", http.MethodPut)
		r.Header.Set("Access-Control-Request-Headers", "X-Forbidden-Header")
		rec := httptest.NewRecorder()
		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("CORS-Enabled"))
		})).ServeHTTP(rec, r)
		res := rec.Result()

		// NOTE: Not sure if it is correct behavior to return 403 here.
		// The preflight request is successful, so you'd expect a 2xx code.
		// As long as the headers are okay, I think it's fine.
		a.So(res.StatusCode, should.Equal, http.StatusForbidden)
		a.So(res.Header.Get("Access-Control-Allow-Origin"), should.BeEmpty)

		body, _ := io.ReadAll(res.Body)
		a.So(body, should.BeEmpty)
	})
}
