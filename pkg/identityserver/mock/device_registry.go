package mockis

import (
	"context"
	"sync"

	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

var errFetchDevice = errors.DefineInternal("fetch_device", "failed to fetch device")

type mockISEndDeviceRegistry struct {
	ttnpb.UnimplementedEndDeviceRegistryServer
	endDevices sync.Map
}

func newEndDeviceRegitry() *mockISEndDeviceRegistry {
	return &mockISEndDeviceRegistry{}
}

func (m *mockISEndDeviceRegistry) Add(ctx context.Context, dev *ttnpb.EndDevice) {
	m.endDevices.Store(unique.ID(ctx, dev.Ids), dev)
}

func (m *mockISEndDeviceRegistry) load(id string) (*ttnpb.EndDevice, error) {
	v, ok := m.endDevices.Load(id)
	if !ok || v == nil {
		return nil, errNotFound.New()
	}
	dev, ok := v.(*ttnpb.EndDevice)
	if !ok {
		return nil, errFetchDevice.New()
	}
	return dev, nil
}

func (m *mockISEndDeviceRegistry) Get(ctx context.Context, in *ttnpb.GetEndDeviceRequest) (*ttnpb.EndDevice, error) {
	return m.load(unique.ID(ctx, in.GetEndDeviceIds()))
}

func (m *mockISEndDeviceRegistry) Update(
	ctx context.Context,
	in *ttnpb.UpdateEndDeviceRequest,
) (*ttnpb.EndDevice, error) {
	dev, err := m.load(unique.ID(ctx, in.GetEndDevice().GetIds()))
	if err != nil {
		return nil, err
	}
	if err := dev.SetFields(in.EndDevice, in.GetFieldMask().GetPaths()...); err != nil {
		return nil, err
	}
	m.Add(ctx, dev)
	return dev, nil
}

type isEndDeviceBatchRegistry struct {
	ttnpb.UnimplementedEndDeviceBatchRegistryServer
	reg *mockISEndDeviceRegistry
}

func newEndDeviceBatchRegitry(devReg *mockISEndDeviceRegistry) *isEndDeviceBatchRegistry {
	return &isEndDeviceBatchRegistry{reg: devReg}
}

// Get implements ttnpb.EndDeviceBatchRegistryServer.
func (m *isEndDeviceBatchRegistry) Get(
	ctx context.Context,
	req *ttnpb.BatchGetEndDevicesRequest,
) (*ttnpb.EndDevices, error) {
	devs := make([]*ttnpb.EndDevice, 0, len(req.DeviceIds))
	for _, id := range req.DeviceIds {
		dev, err := m.reg.load(unique.ID(ctx, &ttnpb.EndDeviceIdentifiers{
			ApplicationIds: req.ApplicationIds,
			DeviceId:       id,
		}))
		if err != nil {
			if errors.IsNotFound(err) {
				continue
			}
			return nil, err
		}
		devs = append(devs, dev)
	}
	return &ttnpb.EndDevices{
		EndDevices: devs,
	}, nil
}
