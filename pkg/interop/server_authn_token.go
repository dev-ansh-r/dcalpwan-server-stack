package interop

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/httpclient"
	"go.thethings.network/lorawan-stack/v3/pkg/packetbroker"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"gopkg.in/square/go-jose.v2/jwt"
)

type packetBrokerTokenVerifier struct {
	publicKeyProvider packetbroker.PublicKeyProvider
	issuer, audience  string
}

func newPacketBrokerTokenVerifier(
	ctx context.Context, issuer, audience string, httpClient httpclient.Provider,
) (tokenVerifier, error) {
	client, err := httpClient.HTTPClient(ctx)
	if err != nil {
		return nil, err
	}
	return &packetBrokerTokenVerifier{
		publicKeyProvider: packetbroker.CachePublicKey(
			packetbroker.PublicKeyFromURL(client, packetbroker.TokenPublicKeysURL(issuer)),
			packetbroker.DefaultPublicKeyCacheTTL,
		),
		issuer:   issuer,
		audience: audience,
	}, nil
}

var errNotPacketBrokerCluster = errors.DefinePermissionDenied("not_packet_broker_cluster",
	"caller not authenticated as Packet Broker cluster",
)

// VerifyNetworkServer verifies the token as Packet Broker cluster token; only Packet Broker clusters can authenticate.
// Packet Broker clusters are authenticated as NetID 000000 and the cluster ID as NSID.
// Packet Broker networks are not allowed to authenticate through LoRaWAN Backend Interfaces.
func (v packetBrokerTokenVerifier) VerifyNetworkServer(
	ctx context.Context, token *jwt.JSONWebToken,
) (*NetworkServerAuthInfo, error) {
	claims, err := packetbroker.Verify(ctx, token, v.publicKeyProvider, v.issuer, v.audience)
	if err != nil {
		return nil, err
	}
	if !claims.PacketBroker.Cluster {
		return nil, errNotPacketBrokerCluster.New()
	}
	return &NetworkServerAuthInfo{
		NetID:     types.NetID{0x0, 0x0, 0x0},
		Addresses: []string{claims.Subject},
	}, nil
}

// VerifyApplicationServer returns an Unauthenticated error; Packet Broker never authenticates as Application Server.
func (packetBrokerTokenVerifier) VerifyApplicationServer(
	context.Context, *jwt.JSONWebToken,
) (*ApplicationServerAuthInfo, error) {
	return nil, errUnauthenticated.New()
}
