package component

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

const testSignalString = "test-signal-string"

type testSignal struct{}

func (ts testSignal) String() string { return testSignalString }
func (ts testSignal) Signal()        {}

func TestRun(t *testing.T) {
	a := assertions.New(t)

	c := MustNew(test.GetLogger(t), &Config{})
	go func() {
		c.terminationSignals <- testSignal{}
	}()
	err := c.Run()
	a.So(err, should.BeNil)
}
