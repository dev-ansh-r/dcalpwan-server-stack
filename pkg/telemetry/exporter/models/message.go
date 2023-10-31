package models

// TelemetryMessage contains all the telemetry data that is to be exported by all services and CLI.
//
// This message is not supposed to be sent in its fullest. Meaning that tasks related to EntitiesCount should send only
// data regarding the entity amount, while the CLI information should be empty.
type TelemetryMessage struct {
	UID           string         `json:"uid"`
	OS            *OSTelemetry   `json:"os,omitempty"`
	CLI           *CLITelemetry  `json:"cli"`
	EntitiesCount *EntitiesCount `json:"entities_count"`
}
