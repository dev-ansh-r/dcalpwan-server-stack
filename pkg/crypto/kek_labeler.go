package crypto

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

// ComponentKEKLabeler provides KEK labels for components.
type ComponentKEKLabeler interface {
	NsKEKLabel(ctx context.Context, netID *types.NetID, addr string) string
	AsKEKLabel(ctx context.Context, addr string) string
}
