package packetbrokeragent

import (
	"context"
	"crypto/tls"

	"go.thethings.network/lorawan-stack/v3/pkg/config/tlsconfig"
	"go.thethings.network/lorawan-stack/v3/pkg/packetbroker"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcclient"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type tlsConfigurator interface {
	GetTLSClientConfig(context.Context, ...tlsconfig.Option) (*tls.Config, error)
}

type authenticator interface {
	AuthInfo(context.Context) (*ttnpb.PacketBrokerNetworkIdentifier, error)
	DialOptions(context.Context) ([]grpc.DialOption, error)
}

type oauth2Authenticator struct {
	tokenSource oauth2.TokenSource
	tlsConfig   tlsConfigurator
}

func newOAuth2(ctx context.Context, oauth2Config OAuth2Config, tlsConfig tlsConfigurator, targetAddresses ...string) authenticator {
	return &oauth2Authenticator{
		tokenSource: packetbroker.TokenSource(ctx, oauth2Config.ClientID, oauth2Config.ClientSecret,
			packetbroker.WithTokenURL(oauth2Config.TokenURL),
			packetbroker.WithScope(packetbroker.ScopeNetworks),
			packetbroker.WithAudienceFromAddresses(targetAddresses...),
		),
		tlsConfig: tlsConfig,
	}
}

func (a *oauth2Authenticator) AuthInfo(ctx context.Context) (*ttnpb.PacketBrokerNetworkIdentifier, error) {
	token, err := a.tokenSource.Token()
	if err != nil {
		return nil, err
	}
	return packetbroker.UnverifiedNetworkIdentifier(token.AccessToken)
}

func (a *oauth2Authenticator) DialOptions(ctx context.Context) ([]grpc.DialOption, error) {
	var (
		tlsConfig *tls.Config
		err       error
	)
	if a.tlsConfig != nil {
		tlsConfig, err = a.tlsConfig.GetTLSClientConfig(ctx)
		if err != nil {
			return nil, err
		}
	}
	res := make([]grpc.DialOption, 2)
	res[0] = grpc.WithPerRPCCredentials(rpcclient.OAuth2(a.tokenSource, tlsConfig == nil))
	if tlsConfig == nil {
		res[1] = grpc.WithInsecure()
	} else {
		res[1] = grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig))
	}
	return res, nil
}
