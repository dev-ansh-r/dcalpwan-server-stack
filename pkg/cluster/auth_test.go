package cluster_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/smarty/assertions"
	clusterauth "go.thethings.network/lorawan-stack/v3/pkg/auth/cluster"
	. "go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc/metadata"
)

func TestVerifySource(t *testing.T) {
	ctx := log.NewContext(test.Context(), test.GetLogger(t))

	a := assertions.New(t)

	key := []byte{0x2A, 0x9C, 0x2C, 0x3C, 0x2A, 0x9C, 0x2A, 0x9C, 0x2A, 0x9C, 0x2A, 0x9C, 0x2A, 0x9C, 0x2A, 0x9C}

	c, err := New(ctx, &Config{
		Keys: []string{
			hex.EncodeToString(key),
		},
	})
	a.So(err, should.BeNil)

	t.Run("empty secret", func(t *testing.T) {
		a := assertions.New(t)

		ctx := c.WithVerifiedSource(ctx)
		a.So(errors.IsUnauthenticated(clusterauth.Authorized(ctx)), should.BeTrue)
	})

	t.Run("invalid secret type", func(t *testing.T) {
		a := assertions.New(t)

		md := metadata.Pairs("authorization", "Basic invalid-secret")
		ctx := metadata.NewIncomingContext(ctx, md)

		ctx = c.WithVerifiedSource(ctx)
		a.So(errors.IsInvalidArgument(clusterauth.Authorized(ctx)), should.BeTrue)
	})

	t.Run("valid secret", func(t *testing.T) {
		a := assertions.New(t)

		md := metadata.Pairs("authorization", fmt.Sprintf("ClusterKey %s", hex.EncodeToString(key)))
		ctx := metadata.NewIncomingContext(ctx, md)

		ctx = c.WithVerifiedSource(ctx)
		a.So(clusterauth.Authorized(ctx), should.BeNil)
	})

	t.Run("wrong secret", func(t *testing.T) {
		a := assertions.New(t)

		md := metadata.Pairs("authorization", "ClusterKey 0102030405060708")
		ctx := metadata.NewIncomingContext(ctx, md)

		ctx = c.WithVerifiedSource(ctx)
		a.So(errors.IsPermissionDenied(clusterauth.Authorized(ctx)), should.BeTrue)
	})
}
