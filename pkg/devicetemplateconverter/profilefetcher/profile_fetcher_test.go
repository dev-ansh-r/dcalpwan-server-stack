package profilefetcher_test

import (
	"context"

	mockdr "go.thethings.network/lorawan-stack/v3/pkg/devicerepository/mock"
	. "go.thethings.network/lorawan-stack/v3/pkg/devicetemplateconverter/profilefetcher"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type mockTemplateFetcher struct {
	mockDR *mockdr.MockDR
}

// GetTemplate makes a request to the Device Repository server with its predefined call options.
func (tf *mockTemplateFetcher) GetTemplate(
	ctx context.Context,
	in *ttnpb.GetTemplateRequest,
) (*ttnpb.EndDeviceTemplate, error) {
	return tf.mockDR.GetTemplate(ctx, in)
}

// MockTemplateFetcher returns an end-device template fetcher that directly points to the DR mock.
func MockTemplateFetcher(mockDR *mockdr.MockDR) TemplateFetcher {
	return &mockTemplateFetcher{
		mockDR: mockDR,
	}
}
