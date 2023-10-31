package mock

import (
	"context"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/scheduling"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Frontend is a mock front-end.
type Frontend struct {
	Up     chan *ttnpb.UplinkMessage
	Status chan *ttnpb.GatewayStatus
	TxAck  chan *ttnpb.TxAcknowledgment
	Down   chan *ttnpb.DownlinkMessage
}

func (*Frontend) Protocol() string                          { return "mock" }
func (*Frontend) SupportsDownlinkClaim() bool               { return true }
func (*Frontend) DutyCycleStyle() scheduling.DutyCycleStyle { return scheduling.DefaultDutyCycleStyle }

// ConnectFrontend connects a new mock front-end to the given server.
// The gateway time starts at Unix epoch.
func ConnectFrontend(ctx context.Context, ids *ttnpb.GatewayIdentifiers, server io.Server) (*Frontend, error) {
	f := &Frontend{
		Up:     make(chan *ttnpb.UplinkMessage, 1),
		Status: make(chan *ttnpb.GatewayStatus, 1),
		TxAck:  make(chan *ttnpb.TxAcknowledgment, 1),
		Down:   make(chan *ttnpb.DownlinkMessage, 1),
	}

	conn, err := server.Connect(ctx, f, ids, &ttnpb.GatewayRemoteAddress{Ip: "127.0.0.1"})
	if err != nil {
		return nil, err
	}
	started := time.Now()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case up := <-f.Up:
				gatewayTime := time.Unix(0, 0).Add(time.Since(started))
				up.ReceivedAt = timestamppb.Now()
				for _, md := range up.RxMetadata {
					md.GpsTime = timestamppb.New(gatewayTime)
				}
				conn.HandleUp(up, nil)
			case status := <-f.Status:
				conn.HandleStatus(status)
			case txAck := <-f.TxAck:
				conn.HandleTxAck(txAck)
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case down := <-conn.Down():
				f.Down <- down
			}
		}
	}()
	return f, nil
}
