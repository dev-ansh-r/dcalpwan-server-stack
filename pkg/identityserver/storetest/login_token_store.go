package storetest

import (
	. "testing"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	is "go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (st *StoreTest) TestLoginTokenStore(t *T) {
	usr1 := st.population.NewUser()

	s, ok := st.PrepareDB(t).(interface {
		Store
		is.LoginTokenStore
	})
	defer st.DestroyDB(t, false)
	if !ok {
		t.Skip("Store does not implement LoginTokenStore")
	}
	defer s.Close()

	start := time.Now().Truncate(time.Second)

	var created *ttnpb.LoginToken

	t.Run("CreateLoginToken", func(t *T) {
		a, ctx := test.New(t)
		var err error
		start := time.Now().Truncate(time.Second)

		created, err = s.CreateLoginToken(ctx, &ttnpb.LoginToken{
			UserIds:   usr1.GetIds(),
			Token:     "TOKEN",
			ExpiresAt: timestamppb.New(start.Add(5 * time.Minute)),
		})

		if a.So(err, should.BeNil) && a.So(created, should.NotBeNil) {
			a.So(created.UserIds, should.Resemble, usr1.GetIds())
			a.So(created.Token, should.Equal, "TOKEN")
			a.So(*ttnpb.StdTime(created.ExpiresAt), should.Equal, start.Add(5*time.Minute))
			a.So(*ttnpb.StdTime(created.CreatedAt), should.HappenWithin, 5*time.Second, start)
			a.So(*ttnpb.StdTime(created.UpdatedAt), should.HappenWithin, 5*time.Second, start)
		}
	})

	t.Run("FindActiveLoginTokens", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.FindActiveLoginTokens(ctx, usr1.GetIds())
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) && a.So(got, should.HaveLength, 1) {
			a.So(got[0], should.Resemble, created)
		}
	})

	t.Run("ConsumeLoginToken", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.ConsumeLoginToken(ctx, "TOKEN")
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
		}
	})

	t.Run("ConsumeLoginToken_Again", func(t *T) {
		a, ctx := test.New(t)
		_, err := s.ConsumeLoginToken(ctx, "TOKEN")
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsFailedPrecondition(err), should.BeTrue)
		}
	})

	t.Run("ConsumeLoginToken_Other", func(t *T) {
		a, ctx := test.New(t)
		_, err := s.ConsumeLoginToken(ctx, "OTHER")
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsNotFound(err), should.BeTrue)
		}
		// TODO: Enable test (https://github.com/TheThingsIndustries/lorawan-stack/issues/3034).
		// _, err = s.ConsumeLoginToken(ctx, "")
		// a.So(err, should.BeNil)
		// if a.So(err, should.NotBeNil) {
		// 	a.So(errors.IsNotFound(err), should.BeTrue)
		// }
	})

	t.Run("FindActiveLoginTokens_AfterConsume", func(t *T) {
		a, ctx := test.New(t)
		got, err := s.FindActiveLoginTokens(ctx, usr1.GetIds())
		if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
			a.So(got, should.BeEmpty)
		}
	})

	t.Run("ConsumeLoginToken_Expired", func(t *T) {
		a, ctx := test.New(t)

		_, err := s.CreateLoginToken(ctx, &ttnpb.LoginToken{
			UserIds:   usr1.GetIds(),
			Token:     "EXPIRED_TOKEN",
			ExpiresAt: timestamppb.New(start.Add(-1 * time.Minute)),
		})
		if !a.So(err, should.BeNil) {
			t.FailNow()
		}

		_, err = s.ConsumeLoginToken(ctx, "EXPIRED_TOKEN")
		if a.So(err, should.NotBeNil) {
			a.So(errors.IsFailedPrecondition(err), should.BeTrue)
		}
	})
}
