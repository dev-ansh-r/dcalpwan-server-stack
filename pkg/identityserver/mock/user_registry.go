package mockis

import (
	"context"
	"sync"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

type mockISUserRegistry struct {
	ttnpb.UnimplementedUserRegistryServer
	ttnpb.UnimplementedUserAccessServer

	users sync.Map
}

func newUserRegistry() *mockISUserRegistry {
	return &mockISUserRegistry{}
}

func (m *mockISUserRegistry) Create(ctx context.Context, req *ttnpb.CreateUserRequest) (*ttnpb.User, error) {
	m.users.Store(unique.ID(ctx, req.User.Ids), req.User)
	return req.User, nil
}

func (m *mockISUserRegistry) Get(ctx context.Context, req *ttnpb.GetUserRequest) (*ttnpb.User, error) {
	loadedUser, ok := m.users.Load(unique.ID(ctx, req.UserIds))
	if !ok {
		return nil, errNotFound
	}
	usr, ok := loadedUser.(*ttnpb.User)
	if !ok {
		panic("stored value is not of type *ttnpb.User")
	}

	return usr, nil
}
