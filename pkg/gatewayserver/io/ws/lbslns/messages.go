package lbslns

import (
	"encoding/json"

	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io/ws/id6"
)

// MessageType is the type of the message.
type MessageType string

// Definition of LoRa Basic Station message types.
const (
	// Upstream types for messages from the Gateway.
	TypeUpstreamVersion              = "version"
	TypeUpstreamJoinRequest          = "jreq"
	TypeUpstreamUplinkDataFrame      = "updf"
	TypeUpstreamProprietaryDataFrame = "propdf"
	TypeUpstreamTxConfirmation       = "dntxed"
	TypeUpstreamTimeSync             = "timesync"
	TypeUpstreamRemoteShell          = "rmtsh"

	// Downstream types for messages from the Network
	TypeDownstreamDownlinkMessage           = "dnmsg"
	TypeDownstreamDownlinkMulticastSchedule = "dnsched"
	TypeDownstreamTimeSync                  = "timesync"
	TypeDownstreamRemoteCommand             = "runcmd"
	TypeDownstreamRemoteShell               = "rmtsh"
)

// DiscoverQuery contains the unique identifier of the gateway.
// This message is sent by the gateway.
type DiscoverQuery struct {
	EUI id6.EUI `json:"router"`
}

// DiscoverResponse contains the response to the discover query.
// This message is sent by the Gateway Server.
type DiscoverResponse struct {
	EUI   id6.EUI `json:"router"`
	Muxs  id6.EUI `json:"muxs,omitempty"`
	URI   string  `json:"uri,omitempty"`
	Error string  `json:"error,omitempty"`
}

// Type returns the message type of the given data.
func Type(data []byte) (string, error) {
	msg := struct {
		Type string `json:"msgtype"`
	}{}
	if err := json.Unmarshal(data, &msg); err != nil {
		return "", err
	}
	return msg.Type, nil
}
