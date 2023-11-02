package mockis

import (
	"context"
	"fmt"
	"sync"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
	"google.golang.org/grpc/metadata"
)

type mockISApplicationRegistry struct {
	ttnpb.UnimplementedApplicationRegistryServer
	ttnpb.UnimplementedApplicationAccessServer

	applications      map[string]*ttnpb.Application
	applicationAuths  map[string][]string
	applicationRights map[string]authKeyToRights

	mu sync.Mutex
}

func newApplicationRegistry() *mockISApplicationRegistry {
	return &mockISApplicationRegistry{
		applications:      make(map[string]*ttnpb.Application),
		applicationAuths:  make(map[string][]string),
		applicationRights: make(map[string]authKeyToRights),
	}
}

func (is *mockISApplicationRegistry) Add(
	ctx context.Context,
	ids *ttnpb.ApplicationIdentifiers,
	key string,
	rights ...ttnpb.Right,
) {
	is.mu.Lock()
	defer is.mu.Unlock()

	uid := unique.ID(ctx, ids)
	is.applications[uid] = &ttnpb.Application{
		Ids: ids,
	}

	var bearerKey string
	if key != "" {
		bearerKey = fmt.Sprintf("Bearer %v", key)
		is.applicationAuths[uid] = append(is.applicationAuths[uid], bearerKey)
	}

	if is.applicationRights[uid] == nil {
		is.applicationRights[uid] = make(authKeyToRights)
	}
	is.applicationRights[uid][bearerKey] = rights
}

func (is *mockISApplicationRegistry) Get(
	ctx context.Context,
	req *ttnpb.GetApplicationRequest,
) (*ttnpb.Application, error) {
	is.mu.Lock()
	defer is.mu.Unlock()

	uid := unique.ID(ctx, req.GetApplicationIds())
	app, ok := is.applications[uid]
	if !ok {
		return nil, errNotFound.New()
	}
	return app, nil
}

func (is *mockISApplicationRegistry) ListRights(
	ctx context.Context,
	ids *ttnpb.ApplicationIdentifiers,
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
	auths, ok := is.applicationAuths[uid]
	if !ok {
		return res, err
	}
	for _, auth := range auths {
		if auth == authorization[0] && is.applicationRights[uid] != nil {
			res.Rights = append(res.Rights, is.applicationRights[uid][auth]...)
		}
	}
	return res, err
}
