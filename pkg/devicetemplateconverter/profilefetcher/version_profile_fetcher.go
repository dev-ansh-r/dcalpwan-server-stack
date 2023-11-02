package profilefetcher

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// versionIDProfileFetcher handles the validation and fetching of end device profile in the Device Repository.
type versionIDProfileFetcher struct{}

// NewFetcherByVersionIDs returns a end-device's profile fetcher that builds its request with vendorIDs.
func NewFetcherByVersionIDs() EndDeviceProfileFetcher {
	return &versionIDProfileFetcher{}
}

// ShouldFetchProfile dictactes if the end device has the necessary fields to fetch its profile.
func (*versionIDProfileFetcher) ShouldFetchProfile(device *ttnpb.EndDevice) bool {
	return device.GetVersionIds().GetBrandId() != "" &&
		device.GetVersionIds().GetModelId() != "" &&
		device.GetVersionIds().GetFirmwareVersion() != "" &&
		device.GetVersionIds().GetBandId() != ""
}

// FetchProfile provides the end device profile.
func (pf *versionIDProfileFetcher) FetchProfile(
	ctx context.Context,
	device *ttnpb.EndDevice,
) (*ttnpb.EndDeviceTemplate, error) {
	versionIDs := &ttnpb.EndDeviceVersionIdentifiers{
		BrandId:         device.GetVersionIds().GetBrandId(),
		ModelId:         device.GetVersionIds().GetModelId(),
		HardwareVersion: device.GetVersionIds().GetHardwareVersion(),
		FirmwareVersion: device.GetVersionIds().GetFirmwareVersion(),
		BandId:          device.GetVersionIds().GetBandId(),
	}
	fetcher, ok := fetcherFromContext(ctx)
	if !ok {
		return nil, nil
	}
	return fetcher.GetTemplate(ctx, &ttnpb.GetTemplateRequest{VersionIds: versionIDs})
}
