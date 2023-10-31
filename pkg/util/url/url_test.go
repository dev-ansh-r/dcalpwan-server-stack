package url_test

import (
	"net/url"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	urlutil "go.thethings.network/lorawan-stack/v3/pkg/util/url"
)

func TestURLClone(t *testing.T) {
	a := assertions.New(t)

	u := &url.URL{
		Scheme: "http",
		Host:   "localhost:1885",
		User:   url.UserPassword("foo", "bar"),
	}
	clone := urlutil.CloneURL(u)
	u.Scheme = "https"
	u.Host = "localhost:8885"
	u.User = url.UserPassword("bar", "foo")

	a.So(u.User, should.NotResemble, clone.User)
	a.So(u, should.NotResemble, clone)
}
