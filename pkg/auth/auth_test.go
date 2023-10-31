package auth_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/auth"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestAuth(t *testing.T) {
	a := assertions.New(t)

	token, err := auth.APIKey.Generate(test.Context(), "")
	a.So(err, should.BeNil)

	tokenType, id, key, err := auth.SplitToken(token)
	a.So(err, should.BeNil)
	a.So(tokenType, should.Equal, auth.APIKey)
	a.So(tokenType.String(), should.Equal, "APIKey")

	a.So(auth.JoinToken(tokenType, id, key), should.Equal, token)

	for _, token := range []string{
		"FOO",             // invalid length
		"FOO.FOO",         // invalid length
		"FOO.FOO.FOO",     // invalid type
		"FOO.FOO.FOO.FOO", // invalid length
	} {
		_, _, _, err := auth.SplitToken(token)
		a.So(err, should.NotBeNil)
	}
}
