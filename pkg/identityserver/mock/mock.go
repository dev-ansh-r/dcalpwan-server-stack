// Package mockis provides a mock structure to the Identity Server.
package mockis

import (
	"context"
	"net"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcserver"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var errNotFound = errors.DefineNotFound("not_found", "not found")

type authKeyToRights map[string][]ttnpb.Right

// MockDefinition contains the structure that is returned by the New(ctx) method of the package, might be used to
// specify IS mock in test cases definitions.
type MockDefinition struct {
	applicationRegistry    *mockISApplicationRegistry
	endDeviceRegistry      *mockISEndDeviceRegistry
	endDeviceBatchRegistry *isEndDeviceBatchRegistry
	entityAccess           *mockEntityAccess
	gatewayRegistry        *mockISGatewayRegistry
	organizationRegistry   *mockISOrganizationRegistry
	userRegistry           *mockISUserRegistry
}

type closeMock func()

// New returns an Identity Server mock along side its address and closing function.
func New(ctx context.Context) (*MockDefinition, string, closeMock) { // nolint:revive
	endDeviceRegistry := newEndDeviceRegitry()

	is := &MockDefinition{
		applicationRegistry:    newApplicationRegistry(),
		endDeviceRegistry:      endDeviceRegistry,
		endDeviceBatchRegistry: newEndDeviceBatchRegitry(endDeviceRegistry),
		entityAccess:           newEntityAccess(),
		gatewayRegistry:        newGatewayRegistry(),
		organizationRegistry:   newOrganizationRegistry(),
		userRegistry:           newUserRegistry(),
	}

	srv := rpcserver.New(ctx)

	ttnpb.RegisterEntityAccessServer(srv.Server, is.entityAccess)

	ttnpb.RegisterApplicationRegistryServer(srv.Server, is.applicationRegistry)
	ttnpb.RegisterApplicationAccessServer(srv.Server, is.applicationRegistry)

	ttnpb.RegisterEndDeviceRegistryServer(srv.Server, is.endDeviceRegistry)
	ttnpb.RegisterEndDeviceBatchRegistryServer(srv.Server, is.endDeviceBatchRegistry)

	ttnpb.RegisterGatewayRegistryServer(srv.Server, is.gatewayRegistry)
	ttnpb.RegisterGatewayAccessServer(srv.Server, is.gatewayRegistry)
	ttnpb.RegisterGatewayBatchAccessServer(srv.Server, is.gatewayRegistry)

	ttnpb.RegisterOrganizationRegistryServer(srv.Server, is.organizationRegistry)
	ttnpb.RegisterOrganizationAccessServer(srv.Server, is.organizationRegistry)

	ttnpb.RegisterUserRegistryServer(srv.Server, is.userRegistry)
	ttnpb.RegisterUserAccessServer(srv.Server, is.userRegistry)

	lis, err := net.Listen("tcp", "")
	if err != nil {
		panic(err)
	}
	go srv.Serve(lis) // nolint:errcheck
	return is, lis.Addr().String(), func() {
		lis.Close()
		srv.GracefulStop()
	}
}

// EndDeviceRegistry returns the methods related to the device registry.
func (m *MockDefinition) EndDeviceRegistry() *mockISEndDeviceRegistry { // nolint:revive
	return m.endDeviceRegistry
}

// ApplicationRegistry returns the methods related to the application registry.
func (m *MockDefinition) ApplicationRegistry() *mockISApplicationRegistry { // nolint:revive
	return m.applicationRegistry
}

// GatewayRegistry returns the methods related to the gateway registry.
func (m *MockDefinition) GatewayRegistry() *mockISGatewayRegistry { // nolint:revive
	return m.gatewayRegistry
}

// EntityAccess returns the methods related to the access entity.
func (m *MockDefinition) EntityAccess() *mockEntityAccess { // nolint:revive
	return m.entityAccess
}

// EndDeviceBatchRegistry returns the methods related to the end device batch registry.
func (m *MockDefinition) EndDeviceBatchRegistry() *isEndDeviceBatchRegistry { // nolint:revive
	return m.endDeviceBatchRegistry
}

// OrganizationRegistry returns the methods related to the organization registry.
func (m *MockDefinition) OrganizationRegistry() *mockISOrganizationRegistry { // nolint:revive
	return m.organizationRegistry
}

// UserRegistry returns the methods related to the user registry.
func (m *MockDefinition) UserRegistry() *mockISUserRegistry { // nolint:revive
	return m.userRegistry
}
