package ttnpb

import "context"

// ContextualEndDevice is an end device with its context.
type ContextualEndDevice struct {
	context.Context
	*EndDevice
}
