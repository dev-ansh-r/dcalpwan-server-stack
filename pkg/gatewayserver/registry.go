package gatewayserver

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// GatewayConnectionStatsRegistry stores, updates and cleans up gateway connection stats.
type GatewayConnectionStatsRegistry interface {
	// Get returns connection stats for a gateway.
	Get(ctx context.Context, ids *ttnpb.GatewayIdentifiers) (*ttnpb.GatewayConnectionStats, error)
	// BatchGet returns connection stats for a batch of gateways.
	BatchGet(
		ctx context.Context,
		ids []*ttnpb.GatewayIdentifiers,
		paths ...string,
	) (map[string]*ttnpb.GatewayConnectionStats, error)
	// Set sets, updates or clears the connection stats for a gateway. Only fields specified in the field mask paths are set.
	Set(
		ctx context.Context,
		ids *ttnpb.GatewayIdentifiers,
		f func(*ttnpb.GatewayConnectionStats) (*ttnpb.GatewayConnectionStats, []string, error),
		ttl time.Duration,
		gets ...string,
	) error
}

// EntityRegistry abstracts the Identity server gateway functions.
type EntityRegistry interface {
	// AssertGatewayRights checks whether the gateway authentication (provied in the context) contains the required rights.
	AssertGatewayRights(ctx context.Context, ids *ttnpb.GatewayIdentifiers, required ...ttnpb.Right) error
	// AssertGatewayBatchRights checks whether the caller has the requested rights on all the requested gateways.
	AssertGatewayBatchRights(ctx context.Context, ids []*ttnpb.GatewayIdentifiers, required ...ttnpb.Right) error
	// Get the identifiers of the gateway that has the given EUI registered.
	GetIdentifiersForEUI(ctx context.Context, in *ttnpb.GetGatewayIdentifiersForEUIRequest) (*ttnpb.GatewayIdentifiers, error)
	// Get the gateway with the given identifiers, selecting the fields specified.
	Get(ctx context.Context, in *ttnpb.GetGatewayRequest) (*ttnpb.Gateway, error)
	// UpdateAntennas updates the gateway antennas.
	UpdateAntennas(ctx context.Context, ids *ttnpb.GatewayIdentifiers, antennas []*ttnpb.GatewayAntenna) error
	// UpdateAttributes updates the gateway attributes. It takes the union of current and new values.
	UpdateAttributes(ctx context.Context, ids *ttnpb.GatewayIdentifiers, current, new map[string]string) error
	// ValidateGatewayID validates the ID of the gateway.
	ValidateGatewayID(ctx context.Context, ids *ttnpb.GatewayIdentifiers) error
}
