// Package provisioning provides a registry and implementations of vendor-specific device provisioners.
package provisioning

import (
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"google.golang.org/protobuf/types/known/structpb"
)

// Provisioner is a device provisioner based on vendor-specific data.
type Provisioner interface {
	// UniqueID returns the vendor-specific unique ID for the given entry.
	UniqueID(entry *structpb.Struct) (string, error)
}

var (
	registry = map[string]Provisioner{}

	errEntry = errors.DefineInvalidArgument("entry", "invalid entry")
)

// Get returns the provisioner by ID.
func Get(id string) Provisioner {
	return registry[id]
}

// Register registers the given provisioner.
// Existing registrations with the same ID will be overwritten.
// This function is not goroutine-safe.
func Register(id string, p Provisioner) {
	registry[id] = p
}
