package udp

import (
	"testing"
	"time"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestTimestamps(t *testing.T) {
	a := assertions.New(t)

	timestamps := newTimestamps(4)

	ret := timestamps.Append(time.Now())
	a.So(ret, should.BeZeroValue)
	for i := 0; i < 3; i++ {
		ret := timestamps.Append(time.Now().Add(time.Hour))
		a.So(ret, should.BeZeroValue)
	}

	val := timestamps.Append(time.Now())
	a.So(val.Before(time.Now()), should.BeTrue)

	val = timestamps.Append(time.Now())
	a.So(val.After(time.Now()), should.BeTrue)
}
