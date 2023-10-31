package rightsutil

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// EventIsVisible returns whether ev is visible given rights in the context.
func EventIsVisible(ctx context.Context, ev events.Event) (bool, error) {
	visibility := ev.Visibility()
	if len(visibility.Rights) == 0 {
		return true, nil
	}
	for _, entityIDs := range ev.Identifiers() {
		switch ids := entityIDs.GetIds().(type) {
		case *ttnpb.EntityIdentifiers_ApplicationIds:
			rights, err := rights.ListApplication(ctx, ids.ApplicationIds)
			if err != nil {
				return false, err
			}
			if len(rights.Implied().Intersect(visibility).GetRights()) > 0 {
				return true, nil
			}
		case *ttnpb.EntityIdentifiers_ClientIds:
			rights, err := rights.ListClient(ctx, ids.ClientIds)
			if err != nil {
				return false, err
			}
			if len(rights.Implied().Intersect(visibility).GetRights()) > 0 {
				return true, nil
			}
		case *ttnpb.EntityIdentifiers_DeviceIds:
			rights, err := rights.ListApplication(ctx, ids.DeviceIds.ApplicationIds)
			if err != nil {
				return false, err
			}
			if len(rights.Implied().Intersect(visibility).GetRights()) > 0 {
				return true, nil
			}
		case *ttnpb.EntityIdentifiers_GatewayIds:
			rights, err := rights.ListGateway(ctx, ids.GatewayIds)
			if err != nil {
				return false, err
			}
			if len(rights.Implied().Intersect(visibility).GetRights()) > 0 {
				return true, nil
			}
		case *ttnpb.EntityIdentifiers_OrganizationIds:
			rights, err := rights.ListOrganization(ctx, ids.OrganizationIds)
			if err != nil {
				return false, err
			}
			if len(rights.Implied().Intersect(visibility).GetRights()) > 0 {
				return true, nil
			}
		case *ttnpb.EntityIdentifiers_UserIds:
			rights, err := rights.ListUser(ctx, ids.UserIds)
			if err != nil {
				return false, err
			}
			if len(rights.Implied().Intersect(visibility).GetRights()) > 0 {
				return true, nil
			}
		}
	}
	return false, nil
}
