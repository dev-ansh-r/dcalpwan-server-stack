package profilefetcher

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

// vendorIDProfileFetcher handles the validation and fetching of end-device profile in the Device Repository.
type vendorIDProfileFetcher struct{}

// NewFetcherByVendorIDs returns a end-device's profile fetcher that builds its request with vendorIDs.
func NewFetcherByVendorIDs() EndDeviceProfileFetcher {
	return &vendorIDProfileFetcher{}
}

// ShouldFetchProfile dictactes if the end-device has the necessary fields to fetch its profile.
func (*vendorIDProfileFetcher) ShouldFetchProfile(device *ttnpb.EndDevice) bool {
	ids := device.GetLoraAllianceProfileIds()
	return ids.GetVendorId() != 0 && ids.GetVendorProfileId() != 0
}

// FetchProfile provides the end-device profile.
func (*vendorIDProfileFetcher) FetchProfile(
	ctx context.Context,
	device *ttnpb.EndDevice,
) (*ttnpb.EndDeviceTemplate, error) {
	profileIdentifiers := &ttnpb.GetTemplateRequest_EndDeviceProfileIdentifiers{
		VendorId:        device.GetLoraAllianceProfileIds().GetVendorId(),
		VendorProfileId: device.GetLoraAllianceProfileIds().GetVendorProfileId(),
	}

	fetcher, ok := fetcherFromContext(ctx)
	if !ok {
		return nil, nil
	}
	return fetcher.GetTemplate(ctx, &ttnpb.GetTemplateRequest{EndDeviceProfileIds: profileIdentifiers})
}
