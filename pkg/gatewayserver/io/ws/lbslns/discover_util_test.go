package lbslns

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/config"
	"go.thethings.network/lorawan-stack/v3/pkg/frequencyplans"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/ratelimit"
	"go.thethings.network/lorawan-stack/v3/pkg/task"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type mockServer struct {
	ids *ttnpb.GatewayIdentifiers
}

func (srv mockServer) GetBaseConfig(ctx context.Context) config.ServiceBase {
	return config.ServiceBase{}
}

func (srv mockServer) FillGatewayContext(ctx context.Context, ids *ttnpb.GatewayIdentifiers) (context.Context, *ttnpb.GatewayIdentifiers, error) {
	return ctx, srv.ids, nil
}

func (mockServer) Connect(
	context.Context, io.Frontend, *ttnpb.GatewayIdentifiers, *ttnpb.GatewayRemoteAddress, ...io.ConnectionOption,
) (*io.Connection, error) {
	return nil, nil
}

func (srv mockServer) GetFrequencyPlans(ctx context.Context, ids *ttnpb.GatewayIdentifiers) (map[string]*frequencyplans.FrequencyPlan, error) {
	return nil, nil
}

func (srv mockServer) ClaimDownlink(ctx context.Context, ids *ttnpb.GatewayIdentifiers) error {
	return nil
}

func (srv mockServer) UnclaimDownlink(ctx context.Context, ids *ttnpb.GatewayIdentifiers) error {
	return nil
}

func (srv mockServer) FromRequestContext(ctx context.Context) context.Context {
	return ctx
}

func (srv mockServer) RateLimiter() ratelimit.Interface {
	return nil
}

func (srv mockServer) ValidateGatewayID(ctx context.Context, ids *ttnpb.GatewayIdentifiers) error {
	return ids.ValidateContext(ctx)
}

func (srv mockServer) StartTask(cfg *task.Config) {
	task.DefaultStartTask(cfg)
}
