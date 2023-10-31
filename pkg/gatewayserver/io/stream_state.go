package io

// MessageStream is a message stream.
type MessageStream uint32

const (
	UplinkStream   MessageStream = iota // UplinkStream is the uplink message stream.
	DownlinkStream                      // DownlinkStream is the downlink message stream.
	TxAckStream                         // TxAckStream is the transmission acknowledgment stream.
	StatusStream                        // StatusStream is the status message stream.
	RTTStream                           // RTTStream is the round-trip times stream.
)

func alwaysOnStreamState(MessageStream) bool { return true }
