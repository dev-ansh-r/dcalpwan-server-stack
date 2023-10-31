package joinserver

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
)

var (
	ErrDevNonceTooSmall  = errDevNonceTooSmall
	ErrNoAppSKey         = errNoAppSKey
	ErrNoFNwkSIntKey     = errNoFNwkSIntKey
	ErrNoNwkSEncKey      = errNoNwkSEncKey
	ErrNoSNwkSIntKey     = errNoSNwkSIntKey
	ErrRegistryOperation = errRegistryOperation
	ErrReuseDevNonce     = errReuseDevNonce
)

type (
	AsJsServer     = asJsServer
	NsJsServer     = nsJsServer
	JsDeviceServer = jsEndDeviceRegistryServer
)

type MockDeviceRegistry struct {
	GetByEUIFunc func(context.Context, types.EUI64, types.EUI64, []string) (*ttnpb.ContextualEndDevice, error)
	GetByIDFunc  func(context.Context, *ttnpb.ApplicationIdentifiers, string, []string) (*ttnpb.EndDevice, error)
	SetByEUIFunc func(
		context.Context,
		types.EUI64,
		types.EUI64,
		[]string,
		func(context.Context, *ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error),
	) (*ttnpb.ContextualEndDevice, error)
	SetByIDFunc func(
		context.Context,
		*ttnpb.ApplicationIdentifiers,
		string,
		[]string, func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error),
	) (*ttnpb.EndDevice, error)
	RangeByIDFunc func(
		context.Context,
		[]string,
		func(context.Context, *ttnpb.EndDeviceIdentifiers, *ttnpb.EndDevice) bool,
	) error
	BatchDeleteFunc func(
		context.Context,
		*ttnpb.ApplicationIdentifiers,
		[]string,
	) ([]*ttnpb.EndDeviceIdentifiers, error)
}

// GetByEUI calls GetByEUIFunc if set and panics otherwise.
func (m MockDeviceRegistry) GetByEUI(ctx context.Context, joinEUI types.EUI64, devEUI types.EUI64, paths []string) (*ttnpb.ContextualEndDevice, error) {
	if m.GetByEUIFunc == nil {
		panic("GetByEUI called, but not set")
	}
	return m.GetByEUIFunc(ctx, joinEUI, devEUI, paths)
}

// GetByID calls GetByIDFunc if set and panics otherwise.
func (m MockDeviceRegistry) GetByID(ctx context.Context, appID *ttnpb.ApplicationIdentifiers, devID string, paths []string) (*ttnpb.EndDevice, error) {
	if m.GetByIDFunc == nil {
		panic("GetByID called, but not set")
	}
	return m.GetByIDFunc(ctx, appID, devID, paths)
}

// SetByEUI calls SetByEUIFunc if set and panics otherwise.
func (m MockDeviceRegistry) SetByEUI(ctx context.Context, joinEUI types.EUI64, devEUI types.EUI64, paths []string, f func(context.Context, *ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.ContextualEndDevice, error) {
	if m.SetByEUIFunc == nil {
		panic("SetByEUI called, but not set")
	}
	return m.SetByEUIFunc(ctx, joinEUI, devEUI, paths, f)
}

// SetByID calls SetByIDFunc if set and panics otherwise.
func (m MockDeviceRegistry) SetByID(ctx context.Context, appID *ttnpb.ApplicationIdentifiers, devID string, paths []string, f func(*ttnpb.EndDevice) (*ttnpb.EndDevice, []string, error)) (*ttnpb.EndDevice, error) {
	if m.SetByIDFunc == nil {
		panic("SetByID called, but not set")
	}
	return m.SetByIDFunc(ctx, appID, devID, paths, f)
}

// SetByID calls SetByIDFunc if set and panics otherwise.
func (m MockDeviceRegistry) RangeByID(ctx context.Context, paths []string, f func(context.Context, *ttnpb.EndDeviceIdentifiers, *ttnpb.EndDevice) bool) error {
	if m.SetByIDFunc == nil {
		panic("SetByID called, but not set")
	}
	return m.RangeByIDFunc(ctx, paths, f)
}

// SetByID calls SetByIDFunc if set and panics otherwise.
func (m MockDeviceRegistry) BatchDelete(
	ctx context.Context,
	appIDs *ttnpb.ApplicationIdentifiers,
	deviceIDs []string,
) ([]*ttnpb.EndDeviceIdentifiers, error) {
	if m.BatchDeleteFunc == nil {
		panic("BatchDelete called, but not set")
	}
	return m.BatchDeleteFunc(ctx, appIDs, deviceIDs)
}

type MockKeyRegistry struct {
	GetByIDFunc func(context.Context, types.EUI64, types.EUI64, []byte, []string) (*ttnpb.SessionKeys, error)
	SetByIDFunc func(
		context.Context,
		types.EUI64,
		types.EUI64,
		[]byte,
		[]string,
		func(*ttnpb.SessionKeys) (*ttnpb.SessionKeys, []string, error),
	) (*ttnpb.SessionKeys, error)
	DeleteFunc      func(context.Context, types.EUI64, types.EUI64) error
	BatchDeleteFunc func(context.Context, []*ttnpb.EndDeviceIdentifiers) error
}

// GetByID calls GetByIDFunc if set and panics otherwise.
func (m MockKeyRegistry) GetByID(ctx context.Context, joinEUI, devEUI types.EUI64, id []byte, paths []string) (*ttnpb.SessionKeys, error) {
	if m.GetByIDFunc == nil {
		panic("GetByID called, but not set")
	}
	return m.GetByIDFunc(ctx, joinEUI, devEUI, id, paths)
}

// SetByID calls SetByIDFunc if set and panics otherwise.
func (m MockKeyRegistry) SetByID(ctx context.Context, joinEUI, devEUI types.EUI64, id []byte, paths []string, f func(*ttnpb.SessionKeys) (*ttnpb.SessionKeys, []string, error)) (*ttnpb.SessionKeys, error) {
	if m.SetByIDFunc == nil {
		panic("SetByID called, but not set")
	}
	return m.SetByIDFunc(ctx, joinEUI, devEUI, id, paths, f)
}

func (m MockKeyRegistry) Delete(ctx context.Context, joinEUI, devEUI types.EUI64) error {
	if m.DeleteFunc == nil {
		panic("Delete called, but not set")
	}
	return m.DeleteFunc(ctx, joinEUI, devEUI)
}

func (m MockKeyRegistry) BatchDelete(ctx context.Context, devIDs []*ttnpb.EndDeviceIdentifiers) error {
	if m.BatchDeleteFunc == nil {
		panic("BatchDelete called, but not set")
	}
	return m.BatchDeleteFunc(ctx, devIDs)
}
