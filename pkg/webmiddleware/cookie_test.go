package webmiddleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/securecookie"
	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	. "go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

func TestCookies(t *testing.T) {
	a := assertions.New(t)

	m := Cookies(
		[]byte("1234123412341234123412341234123412341234123412341234123412341234"),
		[]byte("12341234123412341234123412341234"),
	)

	t.Run("Set and Get Cookie", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodPut, "/", nil)
		rec := httptest.NewRecorder()

		var sc *securecookie.SecureCookie

		m(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sc, _ = GetSecureCookie(r.Context())
			value := map[string]string{
				"foo": "bar",
			}
			if encoded, err := sc.Encode("cookie-name", value); err == nil {
				cookie := &http.Cookie{
					Name:     "cookie-name",
					Value:    encoded,
					Path:     "/",
					Secure:   true,
					HttpOnly: true,
				}
				http.SetCookie(rec, cookie)
			}
		})).ServeHTTP(rec, r)
		res := rec.Result()

		a.So(res.StatusCode, should.Equal, http.StatusOK)
		a.So(res.Header.Get("Set-Cookie"), should.ContainSubstring, "cookie-name")

		cookies := res.Cookies()

		a.So(cookies, should.HaveLength, 1)
		a.So(cookies[0].Name, should.Equal, "cookie-name")
		a.So(cookies[0].Value, should.NotBeEmpty)

		for _, cookie := range res.Cookies() {
			if cookie.Name == "cookie-name" {
				value := make(map[string]string)
				err := sc.Decode("cookie-name", cookie.Value, &value)
				a.So(err, should.BeNil)
				a.So(value["foo"], should.Equal, "bar")
			}
		}
	})
}
