// Package telemetry contains the telemetry configuration and the hash generator that is used in other telemetry related
// packages.
package telemetry

import "time"

// CLI contains information regarding the telemetry collection for the CLI.
type CLI struct {
	Enable bool   `name:"enable" description:"Enables telemetry for CLI"`
	Target string `name:"target" description:"Target to which the information will be sent to"`
}

// EntityCountTelemetry contains information regarding the telemetry collection for the amount of entities.
type EntityCountTelemetry struct {
	Enable   bool          `name:"enable" description:"Enables entity count collection"`
	Interval time.Duration `name:"interval" description:"Interval between each run of the collection"`
}

// Config contains information regarding the telemetry collection.
type Config struct {
	Enable               bool                 `name:"enable" description:"Enables telemetry collection"`
	Target               string               `name:"target" description:"Target to which the information will be sent to"`                                 // nolint:lll
	UIDElements          []string             `name:"uid-elements" description:"Elements that will be used to generate the UID"`                            // nolint:lll
	NumConsumers         uint64               `name:"num-consumers" description:"Number of consumers that will be used to monitor telemetry related tasks"` // nolint:lll
	EntityCountTelemetry EntityCountTelemetry `name:"entity-count-telemetry"`
}
