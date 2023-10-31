package rpcmetadata_test

import (
	"testing"

	"github.com/smarty/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	. "go.thethings.network/lorawan-stack/v3/pkg/rpcmetadata"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestRequestMetadata(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	md1 := MD{
		ID:             "some-id",
		AuthType:       "Key",
		AuthValue:      "foo",
		ServiceType:    "component",
		ServiceVersion: "1.2.3-dev",
		NetAddress:     "localhost",
		Host:           "hostfoo",
		URI:            "fooURI",
	}
	md2 := FromIncomingContext(ctx)

	{
		md, err := md1.GetRequestMetadata(test.Context())
		a.So(err, should.BeNil)
		a.So(md, should.Resemble, map[string]string{
			"id":            "some-id",
			"authorization": "Key foo",
		})
	}

	{
		md, err := md2.GetRequestMetadata(test.Context())
		a.So(err, should.BeNil)
		a.So(md, should.BeEmpty)
	}

	{
		ctx := metadata.NewIncomingContext(test.Context(), metadata.New(map[string]string{
			"id":            "some-id",
			"authorization": "Key foo",
			"host":          "test.local",
		}))
		callOpt, err := WithForwardedAuth(ctx, true)
		a.So(err, should.BeNil)
		requestMD, err := callOpt.(grpc.PerRPCCredsCallOption).Creds.GetRequestMetadata(ctx)
		a.So(err, should.BeNil)
		a.So(requestMD, should.Resemble, map[string]string{
			"id":            "some-id",
			"authorization": "Key foo",
		})
	}

	{
		ctx := metadata.NewIncomingContext(test.Context(), metadata.New(map[string]string{
			"id":   "some-id",
			"host": "test.local",
		}))
		_, err := WithForwardedAuth(ctx, true)
		a.So(errors.IsUnauthenticated(err), should.BeTrue)
	}
}
