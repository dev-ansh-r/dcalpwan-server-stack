// Package mock provides mock implementation of necessary NS interfaces for testing.
package mock

import (
	"context"
	"net"

	"go.thethings.network/lorawan-stack/v3/pkg/rpcserver"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// NS is a mock NS for GS tests.
type NS struct {
	ttnpb.UnimplementedGsNsServer

	upCh    chan *ttnpb.UplinkMessage
	txAckCh chan *ttnpb.GatewayTxAcknowledgment
}

// StartNS starts the mock NS.
func StartNS(ctx context.Context) (*NS, string) {
	ns := &NS{
		upCh:    make(chan *ttnpb.UplinkMessage, 1),
		txAckCh: make(chan *ttnpb.GatewayTxAcknowledgment, 1),
	}
	srv := rpcserver.New(ctx)
	ttnpb.RegisterGsNsServer(srv.Server, ns)
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	go srv.Serve(lis)
	return ns, lis.Addr().String()
}

// HandleUplink implements ttnpb.GsNsServer
func (ns *NS) HandleUplink(ctx context.Context, msg *ttnpb.UplinkMessage) (*emptypb.Empty, error) {
	ns.upCh <- msg
	return ttnpb.Empty, nil
}

// ReportTxAcknowledgment implements ttnpb.GsNsServer
func (ns *NS) ReportTxAcknowledgment(ctx context.Context, msg *ttnpb.GatewayTxAcknowledgment) (*emptypb.Empty, error) {
	ns.txAckCh <- msg
	return ttnpb.Empty, nil
}

// Up returns the upstream channel.
func (ns *NS) Up() <-chan *ttnpb.UplinkMessage {
	return ns.upCh
}

// TxAck returns the TxAck channel.
func (ns *NS) TxAck() <-chan *ttnpb.GatewayTxAcknowledgment {
	return ns.txAckCh
}
