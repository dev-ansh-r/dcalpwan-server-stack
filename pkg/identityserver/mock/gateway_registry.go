package mockis

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmetadata"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

var errNoGatewayRights = errors.DefinePermissionDenied("no_gateway_rights", "no gateway rights")

// DefaultGateway generates a gateway with values that is adequate for most test cases.
func DefaultGateway(ids *ttnpb.GatewayIdentifiers, locationPublic, updateLocationFromStatus bool) *ttnpb.Gateway {
	return &ttnpb.Gateway{
		Ids:              ids,
		FrequencyPlanId:  test.EUFrequencyPlanID,
		FrequencyPlanIds: []string{test.EUFrequencyPlanID},
		Antennas: []*ttnpb.GatewayAntenna{
			{
				Location: &ttnpb.Location{
					Source: ttnpb.LocationSource_SOURCE_REGISTRY,
				},
			},
		},
		GatewayServerAddress:     "mockgatewayserver",
		LocationPublic:           locationPublic,
		UpdateLocationFromStatus: updateLocationFromStatus,
	}
}

type mockISGatewayRegistry struct {
	ttnpb.UnimplementedGatewayRegistryServer
	ttnpb.UnimplementedGatewayAccessServer
	ttnpb.UnimplementedGatewayBatchAccessServer
	gateways      map[string]*ttnpb.Gateway
	gatewayAuths  map[string][]string
	gatewayRights map[string]authKeyToRights

	registeredGateway *ttnpb.GatewayIdentifiers

	mu sync.Mutex
}

func newGatewayRegistry() *mockISGatewayRegistry {
	return &mockISGatewayRegistry{
		gateways:      make(map[string]*ttnpb.Gateway),
		gatewayAuths:  make(map[string][]string),
		gatewayRights: make(map[string]authKeyToRights),
	}
}

func (is *mockISGatewayRegistry) SetRegisteredGateway(gtwIDs *ttnpb.GatewayIdentifiers) {
	is.mu.Lock()
	defer is.mu.Unlock()

	is.registeredGateway = gtwIDs
}

func (is *mockISGatewayRegistry) Add(
	ctx context.Context,
	ids *ttnpb.GatewayIdentifiers,
	key string,
	gtw *ttnpb.Gateway,
	rights ...ttnpb.Right,
) {
	is.mu.Lock()
	defer is.mu.Unlock()

	uid := unique.ID(ctx, ids)
	is.gateways[uid] = gtw

	var bearerKey string
	if key != "" {
		bearerKey = fmt.Sprintf("Bearer %v", key)
		is.gatewayAuths[uid] = []string{bearerKey}
	}
	if is.gatewayRights[uid] == nil {
		is.gatewayRights[uid] = make(authKeyToRights)
	}
	is.gatewayRights[uid][bearerKey] = rights
}

func (is *mockISGatewayRegistry) Get(ctx context.Context, req *ttnpb.GetGatewayRequest) (*ttnpb.Gateway, error) {
	is.mu.Lock()
	defer is.mu.Unlock()

	uid := unique.ID(ctx, req.GetGatewayIds())
	gtw, ok := is.gateways[uid]
	if !ok {
		return nil, errNotFound.New()
	}
	if gtw == nil {
		return nil, errNoGatewayRights.New()
	}
	return gtw, nil
}

func (is *mockISGatewayRegistry) Update(ctx context.Context, req *ttnpb.UpdateGatewayRequest) (*ttnpb.Gateway, error) {
	is.mu.Lock()
	defer is.mu.Unlock()

	uid := unique.ID(ctx, req.Gateway.GetIds())
	gtw, ok := is.gateways[uid]
	if !ok {
		return nil, errNotFound.New()
	}
	if err := gtw.SetFields(req.Gateway, req.FieldMask.GetPaths()...); err != nil {
		return nil, err
	}
	return gtw, nil
}

func (is *mockISGatewayRegistry) Delete(ctx context.Context, ids *ttnpb.GatewayIdentifiers) (*emptypb.Empty, error) {
	is.mu.Lock()
	defer is.mu.Unlock()

	uid := unique.ID(ctx, ids)
	_, ok := is.gateways[uid]
	if !ok {
		return nil, errNotFound.New()
	}
	is.gateways[uid] = nil
	return ttnpb.Empty, nil
}

func (is *mockISGatewayRegistry) ListRights(
	ctx context.Context,
	ids *ttnpb.GatewayIdentifiers,
) (res *ttnpb.Rights, err error) {
	is.mu.Lock()
	defer is.mu.Unlock()

	res = &ttnpb.Rights{}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return res, err
	}
	authorization, ok := md["authorization"]
	if !ok || len(authorization) == 0 {
		return res, err
	}

	uid := unique.ID(ctx, ids)
	auths, ok := is.gatewayAuths[uid]
	if !ok {
		return res, err
	}
	for _, auth := range auths {
		if auth == authorization[0] && is.gatewayRights[uid] != nil {
			res.Rights = append(res.Rights, is.gatewayRights[uid][auth]...)
		}
	}
	return res, err
}

// AssertRights implements GatewayBatchAccess.
func (is *mockISGatewayRegistry) AssertRights(
	ctx context.Context,
	req *ttnpb.AssertGatewayRightsRequest,
) (*emptypb.Empty, error) {
	is.mu.Lock()
	defer is.mu.Unlock()

	md := rpcmetadata.FromIncomingContext(ctx)
	if md.AuthType == "" {
		return nil, errNoGatewayRights.New()
	}
	if strings.ToLower(md.AuthType) != "bearer" {
		return nil, errNoGatewayRights.New()
	}

	for _, ids := range req.GatewayIds {
		uid := unique.ID(ctx, ids)
		auths, ok := is.gatewayAuths[uid]
		if !ok {
			return nil, errNoGatewayRights.New()
		}
		for _, a := range auths {
			if a != fmt.Sprintf("Bearer %s", md.AuthValue) {
				return nil, errNoGatewayRights.New()
			}
			authKeyToRights := is.gatewayRights[uid]
			if authKeyToRights == nil {
				return nil, errNoGatewayRights.New()
			}
			r, ok := authKeyToRights[a]
			if !ok {
				return nil, errNoGatewayRights.New()
			}
			rights := &ttnpb.Rights{Rights: r}
			if len(rights.Intersect(req.Required).GetRights()) == 0 {
				return nil, errNoGatewayRights.New()
			}
		}
	}
	return ttnpb.Empty, nil
}

func (is *mockISGatewayRegistry) GetIdentifiersForEUI(
	_ context.Context,
	req *ttnpb.GetGatewayIdentifiersForEUIRequest,
) (*ttnpb.GatewayIdentifiers, error) {
	is.mu.Lock()
	defer is.mu.Unlock()

	if is.registeredGateway == nil {
		return nil, errNotFound.New()
	}
	if types.MustEUI64(req.Eui).OrZero().Equal(types.MustEUI64(is.registeredGateway.Eui).OrZero()) {
		return is.registeredGateway, nil
	}
	return nil, errNotFound.New()
}
