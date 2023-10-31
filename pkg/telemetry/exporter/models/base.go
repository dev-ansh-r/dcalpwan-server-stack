// Package models contains telemetry models.
package models

// CLITelemetry contains telemetry information about the CLI.
type CLITelemetry struct{}

// OSTelemetry contains telemetry information about the operating system.
type OSTelemetry struct {
	OperatingSystem string `json:"operating_system"`
	Arch            string `json:"arch"`
	BinaryVersion   string `json:"binary_version" `
	GolangVersion   string `json:"golang_version"`
}
