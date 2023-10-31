package identityserver

import (
	"context"
	"strconv"

	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/validate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

var (
	getPaths    = []string{"ids", "created_at", "updated_at"}
	updatePaths = []string{"updated_at"}
)

func cleanFieldMaskPaths(allowedPaths []string, requestedPaths *fieldmaskpb.FieldMask, addPaths, removePaths []string) *fieldmaskpb.FieldMask {
	selected := make(map[string]struct{})
	for _, path := range addPaths {
		selected[path] = struct{}{}
	}
	for _, path := range requestedPaths.GetPaths() {
		selected[path] = struct{}{}
	}
	for _, path := range removePaths {
		delete(selected, path)
	}
	out := make([]string, 0, len(selected))
	for _, path := range allowedPaths {
		if _, ok := selected[path]; ok {
			out = append(out, path)
		}
	}
	return ttnpb.FieldMask(out...)
}

func cleanContactInfo(info []*ttnpb.ContactInfo) {
	for _, info := range info {
		info.ValidatedAt = nil
	}
}

// TODO: Move this logic to validators in API boundary (https://github.com/TheThingsNetwork/lorawan-stack/issues/69).
func validateContactInfo(info []*ttnpb.ContactInfo) error {
	for _, info := range info {
		if info.ContactMethod != ttnpb.ContactMethod_CONTACT_METHOD_EMAIL {
			continue
		}
		if err := validate.Email(info.Value); err != nil {
			return err
		}
	}
	return nil
}

func setTotalHeader(ctx context.Context, total uint64) {
	grpc.SetHeader(ctx, metadata.Pairs("x-total-count", strconv.FormatUint(total, 10)))
}

func idStrings(entityIDs ...*ttnpb.EntityIdentifiers) []string {
	idStrings := make([]string, len(entityIDs))
	for i, entityID := range entityIDs {
		idStrings[i] = entityID.IDString()
	}
	return idStrings
}

// isLastAdmin determines if the userID provided belongs to the last admin of the tenant. Returns an error if the
// condition is true.
func isLastAdmin(ctx context.Context, st store.Store, id *ttnpb.UserIdentifiers) error {
	adminList, err := st.ListAdmins(ctx, store.FieldMask{"ids"})
	if err != nil {
		return err
	}
	if len(adminList) == 1 && proto.Equal(adminList[0].Ids, id) {
		return store.ErrLastAdmin.New()
	}
	return nil
}
