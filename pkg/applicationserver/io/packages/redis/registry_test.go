package redis

import (
	"testing"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/protobuf/types/known/structpb"
)

var appIDs = &ttnpb.ApplicationIdentifiers{
	ApplicationId: "test-app-id",
}

var devIDs = &ttnpb.EndDeviceIdentifiers{
	ApplicationIds: appIDs,
	DeviceId:       "test-dev-id",
}

func TestPkgRegistryClearDefaultAssociations(t *testing.T) {
	t.Parallel()
	a, ctx := test.New(t)
	redisCl, cleanup := test.NewRedis(ctx, "assoc_test")
	t.Cleanup(func() {
		cleanup()
		if err := redisCl.Close(); err != nil {
			t.FailNow()
		}
	})

	registry, err := NewApplicationPackagesRegistry(ctx, redisCl, 10*time.Second)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	appPkgIds := &ttnpb.ApplicationPackageDefaultAssociationIdentifiers{
		ApplicationIds: appIDs,
		FPort:          201,
	}
	expected := &ttnpb.ApplicationPackageDefaultAssociation{
		Ids:         appPkgIds,
		PackageName: "lora-cloud-geolocation-v3",
		Data: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"key": {
					Kind: &structpb.Value_StringValue{
						StringValue: "value",
					},
				},
			},
		},
	}
	_, err = registry.SetDefaultAssociation(ctx, appPkgIds, nil,
		func(apa *ttnpb.ApplicationPackageDefaultAssociation) (
			*ttnpb.ApplicationPackageDefaultAssociation, []string, error,
		) {
			return expected, []string{"ids", "package_name", "data"}, nil
		},
	)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	actual, err := registry.ListDefaultAssociations(ctx, appIDs, []string{"data"})
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	a.So(len(actual), should.Equal, 1)
	a.So(actual[0], should.Resemble, expected)

	err = registry.ClearDefaultAssociations(ctx, appIDs)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	actual, err = registry.ListDefaultAssociations(ctx, appIDs, nil)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	a.So(len(actual), should.Equal, 0)
}

func TestPackageClearAssociations(t *testing.T) {
	t.Parallel()
	a, ctx := test.New(t)
	redisCl, cleanup := test.NewRedis(ctx, "assoc_test")
	t.Cleanup(func() {
		cleanup()
		if err := redisCl.Close(); err != nil {
			t.FailNow()
		}
	})

	registry, err := NewApplicationPackagesRegistry(ctx, redisCl, 10*time.Second)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	appPkgIds := &ttnpb.ApplicationPackageAssociationIdentifiers{
		EndDeviceIds: devIDs,
		FPort:        201,
	}
	expected := &ttnpb.ApplicationPackageAssociation{
		Ids:         appPkgIds,
		PackageName: "lora-cloud-geolocation-v3",
		Data: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"key": {
					Kind: &structpb.Value_StringValue{
						StringValue: "value",
					},
				},
			},
		},
	}
	_, err = registry.SetAssociation(ctx, appPkgIds, nil,
		func(apa *ttnpb.ApplicationPackageAssociation) (
			*ttnpb.ApplicationPackageAssociation, []string, error,
		) {
			return expected, []string{"ids", "package_name", "data"}, nil
		},
	)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	actual, err := registry.ListAssociations(ctx, devIDs, []string{"data"})
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	a.So(len(actual), should.Equal, 1)
	a.So(actual[0], should.Resemble, expected)

	err = registry.ClearAssociations(ctx, devIDs)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	actual, err = registry.ListAssociations(ctx, devIDs, nil)
	if !a.So(err, should.BeNil) {
		t.FailNow()
	}
	a.So(len(actual), should.Equal, 0)
}
