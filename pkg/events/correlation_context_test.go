package events_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestCorrelationContext(t *testing.T) {
	t.Parallel()
	a := assertions.New(t)

	ctx := test.Context()
	correlationIDs := events.CorrelationIDsFromContext(ctx)
	a.So(correlationIDs, should.BeEmpty)

	// Add correlation ID:
	ctx = events.ContextWithCorrelationID(ctx, "foo")
	correlationIDs = events.CorrelationIDsFromContext(ctx)
	a.So(correlationIDs, should.Resemble, []string{"foo"})

	// Add different correlation ID:
	ctx = events.ContextWithCorrelationID(ctx, "baz")
	correlationIDs = events.CorrelationIDsFromContext(ctx)
	a.So(correlationIDs, should.Resemble, []string{"baz", "foo"})

	// Only add new correlation IDs:
	ctx = events.ContextWithCorrelationID(ctx, "bar", "foo")
	correlationIDs = events.CorrelationIDsFromContext(ctx)
	a.So(correlationIDs, should.Resemble, []string{"bar", "baz", "foo"})

	// Do not mutate the passed slice
	base := []string{
		"b",
		"a",
		"c",
	}
	ctx = events.ContextWithCorrelationID(ctx, base...)
	correlationIDs = events.CorrelationIDsFromContext(ctx)
	a.So(correlationIDs, should.Resemble, []string{"a", "b", "bar", "baz", "c", "foo"})
	a.So(base, should.Resemble, []string{
		"b",
		"a",
		"c",
	})
}
