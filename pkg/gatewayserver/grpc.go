package gatewayserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

// GetGatewayConnectionStats returns statistics about a gateway connection.
func (gs *GatewayServer) GetGatewayConnectionStats(ctx context.Context, ids *ttnpb.GatewayIdentifiers) (*ttnpb.GatewayConnectionStats, error) {
	if err := gs.entityRegistry.AssertGatewayRights(ctx, ids, ttnpb.Right_RIGHT_GATEWAY_STATUS_READ); err != nil {
		return nil, err
	}

	uid := unique.ID(ctx, ids)
	if gs.statsRegistry != nil {
		stats, err := gs.statsRegistry.Get(ctx, ids)
		if err != nil || stats == nil {
			if errors.IsNotFound(err) {
				return nil, errNotConnected.WithAttributes("gateway_uid", uid).WithCause(err)
			}
			return nil, err
		}

		return stats, nil
	}

	val, ok := gs.connections.Load(uid)
	if !ok {
		return nil, errNotConnected.WithAttributes("gateway_uid", uid)
	}
	stats, _ := val.(connectionEntry).Stats()
	return stats, nil
}

func applyGatewayConnectionStatsFieldMask(
	dst, src *ttnpb.GatewayConnectionStats,
	paths ...string,
) (*ttnpb.GatewayConnectionStats, error) {
	if dst == nil {
		dst = &ttnpb.GatewayConnectionStats{}
	}
	return dst, dst.SetFields(src, paths...)
}

// BatchGetGatewayConnectionStats implements Gs.
func (gs *GatewayServer) BatchGetGatewayConnectionStats(
	ctx context.Context,
	req *ttnpb.BatchGetGatewayConnectionStatsRequest,
) (*ttnpb.BatchGetGatewayConnectionStatsResponse, error) {
	if err := gs.entityRegistry.AssertGatewayBatchRights(
		ctx,
		req.GatewayIds,
		ttnpb.Right_RIGHT_GATEWAY_STATUS_READ,
	); err != nil {
		return nil, err
	}

	if gs.statsRegistry != nil {
		entries, err := gs.statsRegistry.BatchGet(ctx, req.GatewayIds, req.FieldMask.GetPaths()...)
		if err != nil {
			return nil, err
		}
		return &ttnpb.BatchGetGatewayConnectionStatsResponse{
			Entries: entries,
		}, nil
	}

	// If there isn't a registry, load the (ephemeral) values stored in the Gateway Server connections.
	entries := make(map[string]*ttnpb.GatewayConnectionStats, len(req.GatewayIds))
	for _, id := range req.GatewayIds {
		uid := unique.ID(ctx, id)
		val, ok := gs.connections.Load(uid)
		if !ok {
			continue
		}
		st, _ := val.(connectionEntry).Stats()
		if len(req.FieldMask.GetPaths()) > 0 {
			selected, err := applyGatewayConnectionStatsFieldMask(nil, st, req.FieldMask.GetPaths()...)
			if err != nil {
				return nil, err
			}
			st = selected
		}
		entries[id.GatewayId] = st
	}
	return &ttnpb.BatchGetGatewayConnectionStatsResponse{
		Entries: entries,
	}, nil
}
