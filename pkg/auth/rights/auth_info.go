package rights

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// AuthInfo lists the authentication info with universal rights, whether the caller is admin and the authentication method.
func AuthInfo(ctx context.Context) (authInfo *ttnpb.AuthInfoResponse, err error) {
	if inCtx, ok := authInfoFromContext(ctx); ok {
		return inCtx, nil
	}
	if inCtx, ok := cacheAuthInfoFromContext(ctx); ok {
		return inCtx, nil
	}
	defer func() {
		if err == nil {
			cacheAuthInfoInContext(ctx, authInfo)
		}
	}()
	fetcher, ok := fetcherFromContext(ctx)
	if !ok {
		panic(errNoFetcher)
	}
	authInfo, err = fetcher.AuthInfo(ctx)
	if err != nil {
		if errors.IsPermissionDenied(err) {
			return nil, nil
		}
		return nil, err
	}
	return authInfo, nil
}

var errUnauthenticated = errors.DefineUnauthenticated("unauthenticated", "unauthenticated")

// RequireAuthentication confirms if the authentication information within a context contains any rights, if so,
// the request is considered to be authenticated.
func RequireAuthentication(ctx context.Context) error {
	log.FromContext(ctx).Debug("Authenticate request")
	authInfo, err := AuthInfo(ctx)
	if err != nil {
		log.FromContext(ctx).WithError(err).Debug("Failed to validate authentication information")
		return errUnauthenticated.WithCause(err)
	}

	if authInfo.GetAccessMethod() == nil && len(authInfo.GetUniversalRights().GetRights()) == 0 {
		return errUnauthenticated.New()
	}
	return nil
}
