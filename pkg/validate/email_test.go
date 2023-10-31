package validate

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestEmail(t *testing.T) {
	a := assertions.New(t)
	a.So(Email("daniel@daniel.me"), should.BeNil)
	a.So(Email("root@localhost"), should.BeNil)
	a.So(Email("rootlocalhost.com"), should.NotBeNil)
}
