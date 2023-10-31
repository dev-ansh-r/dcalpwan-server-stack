package fetch_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestHTTP(t *testing.T) {
	a := assertions.New(t)

	// Invalid path
	{
		fetcher, err := fetch.FromHTTP(http.DefaultClient, "")
		if a.So(err, should.BeNil) && a.So(fetcher, should.NotBeNil) {
			_, err = fetcher.File("test")
			a.So(err, should.BeError)
		}

		fetcher, err = fetch.FromHTTP(http.DefaultClient, "/asd")
		if a.So(err, should.NotBeNil) {
			a.So(fetcher, should.BeNil)
		}
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	serverHost := "http://server"
	content := "server content"

	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/success", serverHost), httpmock.NewStringResponder(200, content))
	httpmock.RegisterResponder("GET", fmt.Sprintf("%s/fail", serverHost), httpmock.NewStringResponder(500, ""))

	fetcher, err := fetch.FromHTTP(http.DefaultClient, serverHost)
	if !a.So(err, should.BeNil) {
		t.Fatalf("Failed to construct HTTP fetcher: %v", err)
	}

	// Valid response code
	{
		receivedContent, err := fetcher.File("success")
		if a.So(err, should.BeNil) {
			a.So(string(receivedContent), should.Equal, content)
		}
	}

	// Internal error response code
	{
		_, err := fetcher.File("fail")
		a.So(err, should.NotBeNil)
	}
}
