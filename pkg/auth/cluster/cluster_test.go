package cluster_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/auth/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc/metadata"
)

func TestAuthorized(t *testing.T) {
	a := assertions.New(t)

	if !a.So(func() { cluster.Authorized(context.Background()) }, should.Panic) {
		t.Fatal("Authorized should fail if there is no authentication data")
	}

	authorizedCtx := cluster.NewContext(context.Background(), nil)
	a.So(cluster.Authorized(authorizedCtx), should.BeNil)
	unauthorizedCtx := cluster.NewContext(context.Background(), errors.New("Unauthorized"))
	a.So(cluster.Authorized(unauthorizedCtx), should.NotBeNil)
}

func TestVerify(t *testing.T) {
	a := assertions.New(t)

	keys := [][]byte{
		{0x00, 0xaa, 0x12},
	}
	ctxWithAuthorization := cluster.VerifySource(context.Background(), keys)
	a.So(cluster.Authorized(ctxWithAuthorization), should.NotBeNil)

	md := metadata.MD{}

	for _, tc := range []struct {
		key     []byte
		success bool
	}{
		{
			key:     keys[0],
			success: true,
		},
		{
			key:     []byte{0x00, 0x00},
			success: false,
		},
	} {
		md["authorization"] = []string{fmt.Sprintf("%s %X", cluster.AuthType, tc.key)}
		ctxWithMetadata := metadata.NewIncomingContext(context.Background(), md)
		ctxWithAuthorization = cluster.VerifySource(ctxWithMetadata, keys)
		if tc.success {
			a.So(cluster.Authorized(ctxWithAuthorization), should.BeNil)
		} else {
			a.So(cluster.Authorized(ctxWithAuthorization), should.NotBeNil)
		}
	}
}

func ExampleAuthorized() {
	// Assume this comes from a hypothetical inter-cluster RPC call.
	var ctx context.Context

	if err := cluster.Authorized(ctx); err != nil {
		// return err
	}
}
