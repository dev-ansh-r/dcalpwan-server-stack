package packetbrokeragent

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

type testAuthenticator struct {
	id          *ttnpb.PacketBrokerNetworkIdentifier
	dialOptions []grpc.DialOption
}

func (a *testAuthenticator) AuthInfo(context.Context) (*ttnpb.PacketBrokerNetworkIdentifier, error) {
	return a.id, nil
}

func (a *testAuthenticator) DialOptions(context.Context) ([]grpc.DialOption, error) {
	return a.dialOptions, nil
}

func WithTestAuthenticator(id *ttnpb.PacketBrokerNetworkIdentifier) Option {
	return func(a *Agent) {
		a.authenticator = &testAuthenticator{
			id: id,
			dialOptions: []grpc.DialOption{
				grpc.WithInsecure(),
			},
		}
	}
}
