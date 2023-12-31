// Package cluster contains cluster authentication-related utilities.
package cluster

import (
	"context"
	"crypto/subtle"
	"encoding/hex"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmetadata"
)

type (
	clusterAuthKeyType        struct{}
	clusterAuthKeyFailureType struct{}
)

var (
	// AuthType used to identify components.
	AuthType = "ClusterKey"

	clusterAuthKey        = clusterAuthKeyType{}
	clusterAuthFailureKey = clusterAuthKeyFailureType{}

	errNoClusterKey        = errors.DefineUnauthenticated("no_cluster_key", "no cluster key auth specified")
	errUnsupportedAuthType = errors.DefineInvalidArgument("auth_type", "cluster auth type `{auth_type}` is not supported")
	errInvalidClusterKey   = errors.DefinePermissionDenied("cluster_key", "invalid cluster key")
)

// NewContext returns a context containing the cluster authentication result.
func NewContext(ctx context.Context, err error) context.Context {
	ctx = context.WithValue(ctx, clusterAuthKey, err == nil)
	ctx = context.WithValue(ctx, clusterAuthFailureKey, err)
	return ctx
}

// VerifySource inspects whether the context contains one of the passed cluster keys,
// and returns a context containing the result.
func VerifySource(ctx context.Context, validKeys [][]byte) context.Context {
	err := verifySource(ctx, validKeys)
	return NewContext(ctx, err)
}

func verifySource(ctx context.Context, validKeys [][]byte) error {
	md := rpcmetadata.FromIncomingContext(ctx)
	switch md.AuthType {
	case AuthType:
	case "":
		return errNoClusterKey.New()
	default:
		return errUnsupportedAuthType.WithAttributes("auth_type", md.AuthType)
	}
	key, err := hex.DecodeString(md.AuthValue)
	if err != nil {
		return errInvalidClusterKey.WithCause(err)
	}
	for _, acceptedKey := range validKeys {
		if subtle.ConstantTimeCompare(acceptedKey, key) == 1 {
			return nil
		}
	}
	return errInvalidClusterKey.New()
}

// Authorized returns whether the context has been identified as a cluster call.
// It panics if it does not inherit from `NewContext`.
func Authorized(ctx context.Context) error {
	ok, isStored := ctx.Value(clusterAuthKey).(bool)
	if !isStored {
		panic("call source not verified; did you register the cluster authentication hook for this call?")
	}
	if ok {
		return nil
	}
	return ctx.Value(clusterAuthFailureKey).(error)
}
