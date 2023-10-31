package io

import (
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/scheduling"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UplinkToken returns an uplink token from the given downlink path.
func UplinkToken(ids *ttnpb.GatewayAntennaIdentifiers, timestamp uint32, concentratorTime scheduling.ConcentratorTime, serverTime time.Time, gatewayTime *time.Time) ([]byte, error) {
	token := ttnpb.UplinkToken{
		Ids:              ids,
		Timestamp:        timestamp,
		ServerTime:       timestamppb.New(serverTime),
		ConcentratorTime: int64(concentratorTime),
		GatewayTime:      ttnpb.ProtoTime(gatewayTime),
	}
	return proto.Marshal(&token)
}

// MustUplinkToken returns an uplink token from the given downlink path.
// This function panics if an error occurs. Use UplinkToken to handle errors.
func MustUplinkToken(ids *ttnpb.GatewayAntennaIdentifiers, timestamp uint32, concentratorTime scheduling.ConcentratorTime, serverTime time.Time, gatewayTime *time.Time) []byte {
	token, err := UplinkToken(ids, timestamp, concentratorTime, serverTime, gatewayTime)
	if err != nil {
		panic(err)
	}
	return token
}

// ParseUplinkToken returns the downlink path from the given uplink token.
func ParseUplinkToken(buf []byte) (*ttnpb.UplinkToken, error) {
	var token ttnpb.UplinkToken
	if err := proto.Unmarshal(buf, &token); err != nil {
		return nil, err
	}
	if err := token.ValidateFields(); err != nil {
		return nil, err
	}
	return &token, nil
}
