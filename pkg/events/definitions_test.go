package events_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestDefinitions(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)
	testEvent := events.Define("test", "Test Event", events.WithVisibility(ttnpb.Right_RIGHT_ALL))
	evt := testEvent.New(test.Context())
	a.So(evt.Name(), should.Equal, "test")
	a.So(evt.Visibility().Rights, should.Contain, ttnpb.Right_RIGHT_ALL)
}
