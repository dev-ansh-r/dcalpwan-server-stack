package cayennelpp

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestParseName(t *testing.T) {
	a := assertions.New(t)

	{
		key, channel, err := parseName("digital_in_8")
		a.So(err, should.BeNil)
		a.So(key, should.Equal, digitalInputKey)
		a.So(channel, should.Equal, 8)
	}

	{
		_, _, err := parseName("digital_in_-1")
		a.So(err, should.NotBeNil)
	}

	{
		_, _, err := parseName("_5")
		a.So(err, should.NotBeNil)
	}

	{
		_, _, err := parseName("test")
		a.So(err, should.NotBeNil)
	}

	{
		_, _, err := parseName("test_wrong")
		a.So(err, should.NotBeNil)
	}
}
