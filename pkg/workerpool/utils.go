package workerpool

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// HandlerFromUplinkHandler converts a static uplink handler to a Handler.
func HandlerFromUplinkHandler(
	handler func(context.Context, *ttnpb.ApplicationUp) error,
) Handler[*ttnpb.ApplicationUp] {
	h := func(ctx context.Context, up *ttnpb.ApplicationUp) {
		if err := handler(ctx, up); err != nil {
			log.FromContext(ctx).WithError(err).Warn("Failed to submit message")
		}
	}
	return h
}
