package packages

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// AssociationRegistry is a registry for application package end device associations.
type AssociationRegistry interface {
	// GetAssociation returns the association by its identifiers.
	GetAssociation(ctx context.Context, ids *ttnpb.ApplicationPackageAssociationIdentifiers, paths []string) (*ttnpb.ApplicationPackageAssociation, error)
	// ListAssociations returns all of the associations of the end device.
	ListAssociations(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers, paths []string) ([]*ttnpb.ApplicationPackageAssociation, error)
	// SetAssociation creates, updates or deletes the association by its identifiers.
	SetAssociation(ctx context.Context, ids *ttnpb.ApplicationPackageAssociationIdentifiers, gets []string, f func(*ttnpb.ApplicationPackageAssociation) (*ttnpb.ApplicationPackageAssociation, []string, error)) (*ttnpb.ApplicationPackageAssociation, error)
	// WithPagination adds the pagination information to the context.
	WithPagination(ctx context.Context, limit, page uint32, total *int64) context.Context
	// ClearAssociations removes all application package associations for a specific end device.
	ClearAssociations(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers) error
}

// DefaultAssociationRegistry is a registry for application package default associations.
type DefaultAssociationRegistry interface {
	// GetDefaultAssociation returns the default association by its identifiers.
	GetDefaultAssociation(ctx context.Context, ids *ttnpb.ApplicationPackageDefaultAssociationIdentifiers, paths []string) (*ttnpb.ApplicationPackageDefaultAssociation, error)
	// ListDefaultAssociation returns all of the default associations of the application.
	ListDefaultAssociations(ctx context.Context, ids *ttnpb.ApplicationIdentifiers, paths []string) ([]*ttnpb.ApplicationPackageDefaultAssociation, error)
	// SetDefaultAssociation creates, updates or deletes the default association by its identifiers.
	SetDefaultAssociation(ctx context.Context, ids *ttnpb.ApplicationPackageDefaultAssociationIdentifiers, gets []string, f func(*ttnpb.ApplicationPackageDefaultAssociation) (*ttnpb.ApplicationPackageDefaultAssociation, []string, error)) (*ttnpb.ApplicationPackageDefaultAssociation, error)
	// WithPagination adds the pagination information to the context.
	WithPagination(ctx context.Context, limit, page uint32, total *int64) context.Context
	// ClearDefaultAssociations removes all application package default associations for a specific application.
	ClearDefaultAssociations(ctx context.Context, ids *ttnpb.ApplicationIdentifiers) error
}

// TransactionRegistry is a registry for application packages transactions.
type TransactionRegistry interface {
	EndDeviceTransaction(ctx context.Context, ids *ttnpb.EndDeviceIdentifiers, fPort uint32, packageName string, fn func(ctx context.Context) error) error
}

// Registry is a registry for application packages.
type Registry interface {
	AssociationRegistry
	DefaultAssociationRegistry
	TransactionRegistry
	// Range ranges over the application packages and calls the appropriate callback function, until false is returned.
	Range(
		ctx context.Context, paths []string,
		devFunc func(context.Context, *ttnpb.EndDeviceIdentifiers, *ttnpb.ApplicationPackageAssociation) bool,
		appFunc func(context.Context, *ttnpb.ApplicationIdentifiers, *ttnpb.ApplicationPackageDefaultAssociation) bool,
	) error
}
