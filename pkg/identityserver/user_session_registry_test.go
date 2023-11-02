package identityserver

import (
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/storetest"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
)

func TestUserSessionsRegistry(t *testing.T) {
	t.Parallel()

	p := &storetest.Population{}

	usr1 := p.NewUser()
	key, _ := p.NewAPIKey(usr1.GetEntityIdentifiers(), ttnpb.Right_RIGHT_ALL)
	creds := rpcCreds(key)
	keyWithoutRights, _ := p.NewAPIKey(usr1.GetEntityIdentifiers())
	credsWithoutRights := rpcCreds(keyWithoutRights)

	a, ctx := test.New(t)

	randomUUID := uuid.NewV4().String()

	testWithIdentityServer(t, func(is *IdentityServer, cc *grpc.ClientConn) {
		reg := ttnpb.NewUserSessionRegistryClient(cc)

		_, err := reg.List(ctx, &ttnpb.ListUserSessionsRequest{
			UserIds: usr1.GetIds(),
		}, credsWithoutRights)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		_, err = reg.Delete(ctx, &ttnpb.UserSessionIdentifiers{
			UserIds:   usr1.GetIds(),
			SessionId: randomUUID,
		}, credsWithoutRights)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsPermissionDenied(err), should.BeTrue)
		}

		sessions, err := reg.List(ctx, &ttnpb.ListUserSessionsRequest{
			UserIds: usr1.GetIds(),
		}, creds)
		if a.So(err, should.BeNil) {
			a.So(sessions.Sessions, should.BeEmpty)
		}

		_, err = reg.Delete(ctx, &ttnpb.UserSessionIdentifiers{
			UserIds:   usr1.GetIds(),
			SessionId: randomUUID,
		}, creds)
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}

		var created *ttnpb.UserSession

		err = is.store.Transact(ctx, func(ctx context.Context, st store.Store) error {
			created, err = st.CreateSession(ctx, &ttnpb.UserSession{
				UserIds:   usr1.GetIds(),
				SessionId: randomUUID,
			})
			return err
		})
		if err != nil {
			t.Fatal(err)
		}

		sessions, err = reg.List(ctx, &ttnpb.ListUserSessionsRequest{
			UserIds: usr1.GetIds(),
		}, creds)
		if a.So(err, should.BeNil) && a.So(sessions, should.NotBeNil) {
			a.So(sessions.Sessions, should.HaveLength, 1)
		}

		_, err = reg.Delete(ctx, &ttnpb.UserSessionIdentifiers{
			UserIds:   usr1.GetIds(),
			SessionId: created.SessionId,
		}, creds)
		a.So(err, should.BeNil)

		sessions, err = reg.List(ctx, &ttnpb.ListUserSessionsRequest{
			UserIds: usr1.GetIds(),
		}, creds)
		if a.So(err, should.BeNil) && a.So(sessions, should.NotBeNil) {
			a.So(sessions.Sessions, should.BeEmpty)
		}
	}, withPrivateTestDatabase(p))
}
