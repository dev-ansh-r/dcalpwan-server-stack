package mockis

import (
	"context"
	"sync"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

type mockISOrganizationRegistry struct {
	ttnpb.UnimplementedOrganizationRegistryServer
	ttnpb.UnimplementedOrganizationAccessServer

	orgs sync.Map
}

func newOrganizationRegistry() *mockISOrganizationRegistry {
	return &mockISOrganizationRegistry{}
}

func (m *mockISOrganizationRegistry) Create(
	ctx context.Context, req *ttnpb.CreateOrganizationRequest,
) (*ttnpb.Organization, error) {
	m.orgs.Store(unique.ID(ctx, req.Organization.Ids), req.Organization)
	return req.Organization, nil
}

func (m *mockISOrganizationRegistry) Get(
	ctx context.Context, req *ttnpb.GetOrganizationRequest,
) (*ttnpb.Organization, error) {
	loadedOrganization, ok := m.orgs.Load(unique.ID(ctx, req.OrganizationIds))
	if !ok {
		return nil, errNotFound
	}
	org, ok := loadedOrganization.(*ttnpb.Organization)
	if !ok {
		panic("stored value is not of type *ttnpb.Organization")
	}

	return org, nil
}
