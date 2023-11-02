package storetest

import (
	"fmt"
	"sort"
	. "testing"
	"time"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (st *StoreTest) TestContactInfoStoreCRUD(t *T) {
	app1 := st.population.NewApplication(nil)
	cli1 := st.population.NewClient(nil)
	gtw1 := st.population.NewGateway(nil)
	org1 := st.population.NewOrganization(nil)
	usr1 := st.population.NewUser()
	usr1.PrimaryEmailAddressValidatedAt = nil

	s, ok := st.PrepareDB(t).(interface {
		Store

		store.ContactInfoStore
	})
	defer st.DestroyDB(t, false)
	if !ok {
		t.Skip("Store does not implement ContactInfoStore")
	}
	defer s.Close()

	for _, ids := range []*ttnpb.EntityIdentifiers{
		app1.GetEntityIdentifiers(),
		cli1.GetEntityIdentifiers(),
		gtw1.GetEntityIdentifiers(),
		org1.GetEntityIdentifiers(),
		usr1.GetEntityIdentifiers(),
	} {
		t.Run(ids.EntityType(), func(t *T) { // nolint:paralleltest
			start := time.Now().Truncate(time.Second)
			info := &ttnpb.ContactInfo{
				ContactType:   ttnpb.ContactType_CONTACT_TYPE_TECHNICAL,
				ContactMethod: ttnpb.ContactMethod_CONTACT_METHOD_EMAIL,
				Value:         "foo@example.com",
				Public:        true,
			}
			if ids.EntityType() == "user" && ids.IDString() == usr1.IDString() {
				info.Value = usr1.PrimaryEmailAddress
			}
			t.Run("SetContactInfo", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				created, err := s.SetContactInfo(ctx, ids, []*ttnpb.ContactInfo{
					info,
				})
				if a.So(err, should.BeNil) && a.So(created, should.NotBeNil) && a.So(created, should.HaveLength, 1) {
					a.So(created[0], should.Resemble, info)
				}
			})

			t.Run("GetContactInfo", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				got, err := s.GetContactInfo(ctx, ids)
				if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) && a.So(got, should.HaveLength, 1) {
					a.So(got[0], should.Resemble, info)
				}
			})

			validationID := fmt.Sprintf("%s_%s_validation", ids.EntityType(), ids.IDString())

			t.Run("CreateValidation_Expired", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				_, err := s.CreateValidation(ctx, &ttnpb.ContactInfoValidation{
					Id:          validationID,
					Token:       "EXPIRED_TOKEN",
					Entity:      ids,
					ContactInfo: []*ttnpb.ContactInfo{info},
					ExpiresAt:   timestamppb.New(start.Add(-1 * time.Minute)),
				})
				if !a.So(err, should.BeNil) {
					t.FailNow()
				}
			})

			// Used for a refresh interval test.
			var createAt time.Time

			t.Run("CreateValidation", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				created, err := s.CreateValidation(ctx, &ttnpb.ContactInfoValidation{
					Id:          validationID,
					Token:       "TOKEN",
					Entity:      ids,
					ContactInfo: []*ttnpb.ContactInfo{info},
					ExpiresAt:   timestamppb.New(start.Add(5 * time.Minute)),
				})
				if a.So(err, should.BeNil) && a.So(created, should.NotBeNil) {
					a.So(created.Id, should.Equal, validationID)
					a.So(created.Token, should.Equal, "TOKEN")
					a.So(created.Entity, should.Resemble, ids)
					a.So(created.ContactInfo, should.Resemble, []*ttnpb.ContactInfo{info})
					a.So(*ttnpb.StdTime(created.ExpiresAt), should.Equal, start.Add(5*time.Minute))
					a.So(*ttnpb.StdTime(created.CreatedAt), should.HappenWithin, 5*time.Second, start)
				}
				createAt = created.CreatedAt.AsTime()
			})

			t.Run("CreateValidation_AfterCreate", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				_, err := s.CreateValidation(ctx, &ttnpb.ContactInfoValidation{
					Id:          validationID,
					Token:       "OTHER_TOKEN",
					Entity:      ids,
					ContactInfo: []*ttnpb.ContactInfo{info},
					ExpiresAt:   timestamppb.New(start.Add(5 * time.Minute)),
				})
				if a.So(err, should.NotBeNil) {
					a.So(errors.IsAlreadyExists(err), should.BeTrue)
				}
			})

			t.Run("ListRefreshableValidations", func(t *T) { // nolint:paralleltest
				t.Run("valid refresh interval", func(t *T) { // nolint:paralleltest
					a, ctx := test.New(t)
					// Any validation from now
					validations, err := s.ListRefreshableValidations(ctx, ids, 0)
					if a.So(err, should.BeNil) && a.So(validations, should.NotBeNil) {
						a.So(validations, should.HaveLength, 1)
						v := validations[0]
						a.So(v.Id, should.Equal, validationID)
						a.So(v.Token, should.Equal, "TOKEN")
						a.So(v.Entity, should.Resemble, ids)
						a.So(v.ContactInfo, should.Resemble, []*ttnpb.ContactInfo{{
							ContactMethod: info.ContactMethod,
							Value:         info.Value,
						}})
						a.So(*ttnpb.StdTime(v.ExpiresAt), should.Equal, start.Add(5*time.Minute))
						a.So(*ttnpb.StdTime(v.CreatedAt), should.HappenWithin, 5*time.Second, start)
						a.So(*ttnpb.StdTime(v.UpdatedAt), should.HappenWithin, 5*time.Second, start)
					}
				})

				t.Run("invalid refresh interval", func(t *T) { // nolint:paralleltest
					a, ctx := test.New(t)
					invalidRefreshInterval := 2 * time.Since(createAt)
					validations, err := s.ListRefreshableValidations(ctx, ids, invalidRefreshInterval)
					if !a.So(err, should.BeNil) || !a.So(validations, should.HaveLength, 0) {
						t.Fail()
					}
				})
			})

			t.Run("FetchExpiredValidation", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				_, err := s.GetValidation(ctx, &ttnpb.ContactInfoValidation{
					Id:    validationID,
					Token: "EXPIRED_TOKEN",
				})
				if a.So(err, should.NotBeNil) {
					a.So(errors.IsFailedPrecondition(err), should.BeTrue)
				}
			})

			t.Run("RefreshValidation", func(t *T) { // nolint:paralleltest
				t.Run("Valid", func(t *T) {
					a, ctx := test.New(t)

					validation := &ttnpb.ContactInfoValidation{
						Id:          validationID,
						Token:       "TOKEN",
						Entity:      ids,
						ContactInfo: []*ttnpb.ContactInfo{info},
					}

					original, err := s.GetValidation(ctx, validation)
					if !a.So(err, should.BeNil) || !a.So(original.GetUpdatedAt(), should.NotBeNil) {
						t.FailNow()
					}

					err = s.RefreshValidation(ctx, original)
					a.So(err, should.BeNil)

					// Get validation again and check if updated_at has changed.
					after, err := s.GetValidation(ctx, validation)
					if !a.So(err, should.BeNil) || !a.So(after.GetUpdatedAt(), should.NotBeNil) {
						t.FailNow()
					}
					a.So(original.UpdatedAt.AsTime().Before(after.UpdatedAt.AsTime()), should.BeTrue)

					// Retry refresh with original validation, as updated_at should have changed it should not refresh.
					err = s.RefreshValidation(ctx, original)
					a.So(errors.IsNotFound(err), should.BeTrue)
				})
				t.Run("Invalid", func(t *T) { // nolint:paralleltest
					a, ctx := test.New(t)
					err := s.RefreshValidation(ctx, &ttnpb.ContactInfoValidation{
						Id:    validationID,
						Token: "INVALID_TOKEN",
					})
					a.So(errors.IsNotFound(err), should.BeTrue)
				})
			})

			t.Run("Validate", func(t *T) { // nolint:paralleltest
				validation := &ttnpb.ContactInfoValidation{
					Id:          validationID,
					Token:       "TOKEN",
					Entity:      ids,
					ContactInfo: []*ttnpb.ContactInfo{info},
				}
				t.Run("GetValidation", func(t *T) { // nolint:paralleltest
					// This test is meant to impact `validation` and must run before any other test within `Validate`.
					a, ctx := test.New(t)
					var err error
					validation, err = s.GetValidation(ctx, validation)
					a.So(err, should.BeNil)
					a.So(validation.Id, should.Equal, validationID)
					a.So(validation.Token, should.Equal, "TOKEN")
					a.So(len(validation.ContactInfo), should.Equal, 1)
				})

				t.Run("ValidateContactInfo", func(t *T) { // nolint:paralleltest
					t.Run("ContactInfoDoesNotExist", func(t *T) { // nolint:paralleltest
						a, ctx := test.New(t)
						v := &ttnpb.ContactInfoValidation{
							Id:          validationID,
							Token:       "TOKEN",
							Entity:      ids,
							ContactInfo: []*ttnpb.ContactInfo{{}}, // Empty contact info.
						}
						err := s.ValidateContactInfo(ctx, v)
						a.So(errors.IsNotFound(err), should.BeTrue)
					})

					t.Run("NoContactInfo", func(t *T) { // nolint:paralleltest
						a, ctx := test.New(t)
						v := &ttnpb.ContactInfoValidation{
							Id:          validationID,
							Token:       "TOKEN",
							Entity:      ids,
							ContactInfo: []*ttnpb.ContactInfo{},
						}
						err := s.ValidateContactInfo(ctx, v)
						a.So(errors.IsFailedPrecondition(err), should.BeTrue)
					})

					t.Run("Valid", func(t *T) { // nolint:paralleltest
						a, ctx := test.New(t)
						err := s.ValidateContactInfo(ctx, validation)
						a.So(err, should.BeNil)
						// Checks if contactInfo was validated.
						got, err := s.GetContactInfo(ctx, ids)
						if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) && a.So(got, should.HaveLength, 1) {
							a.So(got[0].ValidatedAt, should.NotBeNil)
						}
					})
				})
				t.Run("ExpireValidation", func(t *T) { // nolint:paralleltest
					a, ctx := test.New(t)
					err := s.ExpireValidation(ctx, validation)
					a.So(err, should.BeNil)
				})
			})
			t.Run("FetchUsedValidation_AfterValidate", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				_, err := s.GetValidation(ctx, &ttnpb.ContactInfoValidation{
					Id:    validationID,
					Token: "TOKEN",
				})
				if a.So(err, should.NotBeNil) {
					a.So(errors.IsFailedPrecondition(err), should.BeTrue)
				}
			})
			t.Run("FetchNonExistentValidation", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				_, err := s.GetValidation(ctx, &ttnpb.ContactInfoValidation{
					Id:    validationID,
					Token: "OTHER_TOKEN",
				})
				if a.So(err, should.NotBeNil) {
					a.So(errors.IsNotFound(err), should.BeTrue)
				}
			})

			t.Run("GetContactInfo_AfterValidate", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				got, err := s.GetContactInfo(ctx, ids)
				if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) && a.So(got, should.HaveLength, 1) {
					a.So(got[0].ValidatedAt, should.NotBeNil)
				}
			})

			if ids.EntityType() == "user" && ids.IDString() == usr1.IDString() {
				t.Run("GetUser_AfterValidate", func(t *T) { // nolint:paralleltest
					a, ctx := test.New(t)
					got, err := s.(store.UserStore).GetUser(ctx, usr1.GetIds(), fieldMask("primary_email_address_validated_at"))
					if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) {
						a.So(got.PrimaryEmailAddressValidatedAt, should.NotBeNil)
					}
				})
			}

			updatedInfo := &ttnpb.ContactInfo{
				ContactType:   ttnpb.ContactType_CONTACT_TYPE_TECHNICAL,
				ContactMethod: ttnpb.ContactMethod_CONTACT_METHOD_EMAIL,
				Value:         "foo@example.com",
				Public:        false,
				ValidatedAt:   timestamppb.New(start.Add(time.Minute)),
			}
			extraInfo := &ttnpb.ContactInfo{
				ContactType:   ttnpb.ContactType_CONTACT_TYPE_TECHNICAL,
				ContactMethod: ttnpb.ContactMethod_CONTACT_METHOD_EMAIL,
				Value:         "extra@example.com",
				Public:        true,
				ValidatedAt:   timestamppb.New(start.Add(time.Minute)),
			}

			t.Run("UpdateContactInfo", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				updated, err := s.SetContactInfo(ctx, ids, []*ttnpb.ContactInfo{
					updatedInfo,
					extraInfo,
				})
				if a.So(err, should.BeNil) && a.So(updated, should.NotBeNil) && a.So(updated, should.HaveLength, 2) {
					sort.Slice(updated, func(i, j int) bool { return updated[i].Value < updated[j].Value })
					a.So(updated[0], should.Resemble, extraInfo)
					a.So(updated[1], should.Resemble, updatedInfo)
				}
			})

			t.Run("GetContactInfo_AfterUpdate", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				got, err := s.GetContactInfo(ctx, ids)
				if a.So(err, should.BeNil) && a.So(got, should.NotBeNil) && a.So(got, should.HaveLength, 2) {
					sort.Slice(got, func(i, j int) bool { return got[i].Value < got[j].Value })
					a.So(got[0], should.Resemble, extraInfo)
					a.So(got[1], should.Resemble, updatedInfo)
				}
			})

			t.Run("DeleteEntityContactInfo", func(t *T) { // nolint:paralleltest
				a, ctx := test.New(t)
				err := s.DeleteEntityContactInfo(ctx, ids)
				a.So(err, should.BeNil)
			})
		})
	}
}
