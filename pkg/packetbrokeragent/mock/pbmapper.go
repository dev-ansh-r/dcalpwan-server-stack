package mock

import (
	"context"
	"testing"

	mappingpb "go.packetbroker.org/api/mapping/v2"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// PBMapper is a mock Packet Broker Mapper.
type PBMapper struct {
	*grpc.Server
	UpdateGatewayHandler func(ctx context.Context, in *mappingpb.UpdateGatewayRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

// NewPBMapper instantiates a new mock Packet Broker Data Plane.
func NewPBMapper(tb testing.TB) *PBMapper {
	mp := &PBMapper{
		Server: grpc.NewServer(
			grpc.UnaryInterceptor(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
				ctx = test.ContextWithTB(ctx, tb)
				return handler(ctx, req)
			}),
		),
	}
	mappingpb.RegisterMapperServer(mp.Server, &pbMapper{PBMapper: mp})
	return mp
}

type pbMapper struct {
	mappingpb.UnimplementedMapperServer

	*PBMapper
}

func (s *pbMapper) UpdateGateway(ctx context.Context, req *mappingpb.UpdateGatewayRequest) (*emptypb.Empty, error) {
	if s.UpdateGatewayHandler == nil {
		panic("UpdateGateway called but not set")
	}
	return s.UpdateGatewayHandler(ctx, req)
}
