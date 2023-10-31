package rights

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

// ListApplication lists the rights for the given application ID in the context.
func ListApplication(ctx context.Context, id *ttnpb.ApplicationIdentifiers) (rights *ttnpb.Rights, err error) {
	uid := unique.ID(ctx, id)
	if inCtx, ok := fromContext(ctx); ok {
		r, _ := inCtx.ApplicationRights.GetRights(uid)
		return r, nil
	}
	if inCtx, ok := cacheFromContext(ctx); ok {
		if r, ok := inCtx.ApplicationRights.GetRights(uid); ok {
			return r, nil
		}
	}
	defer func() {
		if err == nil {
			cacheInContext(ctx, func(r *Rights) { r.setApplicationRights(uid, rights) })
		}
	}()
	fetcher, ok := fetcherFromContext(ctx)
	if !ok {
		panic(errNoFetcher)
	}
	rights, err = fetcher.ApplicationRights(ctx, id)
	if err != nil {
		if errors.IsPermissionDenied(err) {
			return &ttnpb.Rights{}, nil
		}
		return nil, err
	}
	return rights, nil
}

// ListClient lists the rights for the given client ID in the context.
func ListClient(ctx context.Context, id *ttnpb.ClientIdentifiers) (rights *ttnpb.Rights, err error) {
	uid := unique.ID(ctx, id)
	if inCtx, ok := fromContext(ctx); ok {
		r, _ := inCtx.ClientRights.GetRights(uid)
		return r, nil
	}
	if inCtx, ok := cacheFromContext(ctx); ok {
		if r, ok := inCtx.ClientRights.GetRights(uid); ok {
			return r, nil
		}
	}
	defer func() {
		if err == nil {
			cacheInContext(ctx, func(r *Rights) { r.setClientRights(uid, rights) })
		}
	}()
	fetcher, ok := fetcherFromContext(ctx)
	if !ok {
		panic(errNoFetcher)
	}
	rights, err = fetcher.ClientRights(ctx, id)
	if err != nil {
		if errors.IsPermissionDenied(err) {
			return &ttnpb.Rights{}, nil
		}
		return nil, err
	}
	return rights, nil
}

// ListGateway lists the rights for the given gateway ID in the context.
func ListGateway(ctx context.Context, id *ttnpb.GatewayIdentifiers) (rights *ttnpb.Rights, err error) {
	uid := unique.ID(ctx, id)
	if inCtx, ok := fromContext(ctx); ok {
		r, _ := inCtx.GatewayRights.GetRights(uid)
		return r, nil
	}
	if inCtx, ok := cacheFromContext(ctx); ok {
		if r, ok := inCtx.GatewayRights.GetRights(uid); ok {
			return r, nil
		}
	}
	defer func() {
		if err == nil {
			cacheInContext(ctx, func(r *Rights) { r.setGatewayRights(uid, rights) })
		}
	}()
	fetcher, ok := fetcherFromContext(ctx)
	if !ok {
		panic(errNoFetcher)
	}
	rights, err = fetcher.GatewayRights(ctx, id)
	if err != nil {
		if errors.IsPermissionDenied(err) {
			return &ttnpb.Rights{}, nil
		}
		return nil, err
	}
	return rights, nil
}

// ListOrganization lists the rights for the given organization ID in the context.
func ListOrganization(ctx context.Context, id *ttnpb.OrganizationIdentifiers) (rights *ttnpb.Rights, err error) {
	uid := unique.ID(ctx, id)
	if inCtx, ok := fromContext(ctx); ok {
		r, _ := inCtx.OrganizationRights.GetRights(uid)
		return r, nil
	}
	if inCtx, ok := cacheFromContext(ctx); ok {
		if r, ok := inCtx.OrganizationRights.GetRights(uid); ok {
			return r, nil
		}
	}
	defer func() {
		if err == nil {
			cacheInContext(ctx, func(r *Rights) { r.setOrganizationRights(uid, rights) })
		}
	}()
	fetcher, ok := fetcherFromContext(ctx)
	if !ok {
		panic(errNoFetcher)
	}
	rights, err = fetcher.OrganizationRights(ctx, id)
	if err != nil {
		if errors.IsPermissionDenied(err) {
			return &ttnpb.Rights{}, nil
		}
		return nil, err
	}
	return rights, nil
}

// ListUser lists the rights for the given user ID in the context.
func ListUser(ctx context.Context, id *ttnpb.UserIdentifiers) (rights *ttnpb.Rights, err error) {
	uid := unique.ID(ctx, id)
	if inCtx, ok := fromContext(ctx); ok {
		r, _ := inCtx.UserRights.GetRights(uid)
		return r, nil
	}
	if inCtx, ok := cacheFromContext(ctx); ok {
		if r, ok := inCtx.UserRights.GetRights(uid); ok {
			return r, nil
		}
	}
	defer func() {
		if err == nil {
			cacheInContext(ctx, func(r *Rights) { r.setUserRights(uid, rights) })
		}
	}()
	fetcher, ok := fetcherFromContext(ctx)
	if !ok {
		panic(errNoFetcher)
	}
	rights, err = fetcher.UserRights(ctx, id)
	if err != nil {
		if errors.IsPermissionDenied(err) {
			return &ttnpb.Rights{}, nil
		}
		return nil, err
	}
	return rights, nil
}
