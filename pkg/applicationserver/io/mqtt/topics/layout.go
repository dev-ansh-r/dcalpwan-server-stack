// Package topics implements MQTT topic layouts.
package topics

// Layout represents an MQTT topic layout.
type Layout interface {
	AcceptedTopic(applicationUID string, requested []string) (accepted []string, ok bool)

	UplinkMessageTopic(applicationUID, deviceID string) []string
	UplinkNormalizedTopic(applicationUID, deviceID string) []string
	JoinAcceptTopic(applicationUID, deviceID string) []string
	DownlinkAckTopic(applicationUID, deviceID string) []string
	DownlinkNackTopic(applicationUID, deviceID string) []string
	DownlinkSentTopic(applicationUID, deviceID string) []string
	DownlinkFailedTopic(applicationUID, deviceID string) []string
	DownlinkQueuedTopic(applicationUID, deviceID string) []string
	DownlinkQueueInvalidatedTopic(applicationUID, deviceID string) []string
	LocationSolvedTopic(applicationUID, deviceID string) []string
	ServiceDataTopic(applicationUID, deviceID string) []string

	DownlinkPushTopic(applicationUID, deviceID string) []string
	IsDownlinkPushTopic(parts []string) bool
	ParseDownlinkPushTopic(parts []string) (deviceID string)
	DownlinkReplaceTopic(applicationUID, deviceID string) []string
	IsDownlinkReplaceTopic(parts []string) bool
	ParseDownlinkReplaceTopic(parts []string) (deviceID string)
}
