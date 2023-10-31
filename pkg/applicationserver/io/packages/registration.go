package packages

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// ApplicationPackageHandler handles upstream traffic from the Application Server.
type ApplicationPackageHandler interface {
	Package() *ttnpb.ApplicationPackage
	HandleUp(context.Context, *ttnpb.ApplicationPackageDefaultAssociation, *ttnpb.ApplicationPackageAssociation, *ttnpb.ApplicationUp) error
}

var errNotImplemented = errors.DefineUnimplemented("package_not_implemented", "package `{name}` is not implemented")
