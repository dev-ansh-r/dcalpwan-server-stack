package packetbrokeragent

import (
	"context"

	clusterauth "go.thethings.network/lorawan-stack/v3/pkg/auth/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type nsPbaServer struct {
	ttnpb.UnimplementedNsPbaServer

	contextDecoupler contextDecoupler
	downstreamCh     chan *downlinkMessage
	frequencyPlans   GetFrequencyPlansStore
}

// PublishDownlink is called by the Network Server when a downlink message needs to get scheduled via Packet Broker.
func (s *nsPbaServer) PublishDownlink(ctx context.Context, down *ttnpb.DownlinkMessage) (*emptypb.Empty, error) {
	if err := clusterauth.Authorized(ctx); err != nil {
		return nil, err
	}

	ctx = events.ContextWithCorrelationID(ctx, down.CorrelationIds...)
	ctx = appendDownlinkCorrelationID(ctx)
	down.CorrelationIds = events.CorrelationIDsFromContext(ctx)

	fps, err := s.frequencyPlans(ctx)
	if err != nil {
		return nil, err
	}

	msg, token, err := toPBDownlink(ctx, down, fps)
	if err != nil {
		log.FromContext(ctx).WithError(err).Warn("Failed to convert outgoing downlink message")
		return nil, err
	}

	ctxMsg := &downlinkMessage{
		Context:                      s.contextDecoupler.FromRequestContext(ctx),
		PacketBrokerAgentUplinkToken: token,
		DownlinkMessage:              msg,
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case s.downstreamCh <- ctxMsg:
		return ttnpb.Empty, nil
	}
}
