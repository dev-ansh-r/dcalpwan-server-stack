package storetest

import (
	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

func fieldMask(paths ...string) store.FieldMask {
	return ttnpb.ExcludeFields(
		paths,
		"ids",
		"created_at",
		"updated_at",
		"deleted_at",
	)
}

var attributes = map[string]string{
	"foo": "bar",
	"bar": "baz",
	"baz": "qux",
}

var updatedAttributes = map[string]string{
	"attribute": "new",
	"foo":       "bar",
	"bar":       "updated",
}
