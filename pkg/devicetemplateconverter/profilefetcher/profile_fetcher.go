// Package profilefetcher contains definitions of end-device's profile fetchers based on the VersionIDs or VendorIDs.
package profilefetcher

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmetadata"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"google.golang.org/grpc"
)

// EndDeviceProfileFetcher specifies the necessary methods to ascertain if the end device contains the necessary
// information to make the profile template request and for performing the request itself.
type EndDeviceProfileFetcher interface {
	ShouldFetchProfile(*ttnpb.EndDevice) bool
	FetchProfile(context.Context, *ttnpb.EndDevice) (*ttnpb.EndDeviceTemplate, error)
}

// TemplateFetcher implements the steps required to obtain a end-device template from the Device Repository.
type TemplateFetcher interface {
	GetTemplate(ctx context.Context, in *ttnpb.GetTemplateRequest) (*ttnpb.EndDeviceTemplate, error)
}

// Component abstracts the underlying *component.Component.
type Component interface {
	GetPeerConn(ctx context.Context, role ttnpb.ClusterRole, ids cluster.EntityIdentifiers) (*grpc.ClientConn, error)
	AllowInsecureForCredentials() bool
}

type templateFetcher struct {
	c Component
}

// NewTemplateFetcher returns an end-device template fetcher with predefined call options.
func NewTemplateFetcher(c Component) TemplateFetcher {
	return &templateFetcher{
		c: c,
	}
}

// GetTemplate makes a request to the Device Repository server with its predefined call options.
func (tf *templateFetcher) GetTemplate(
	ctx context.Context,
	in *ttnpb.GetTemplateRequest,
) (*ttnpb.EndDeviceTemplate, error) {
	conn, err := tf.c.GetPeerConn(ctx, ttnpb.ClusterRole_DEVICE_REPOSITORY, nil)
	if err != nil {
		log.FromContext(ctx).WithError(err).Warn("Failed to get Device Repository peer")
		return nil, err
	}

	opt, err := rpcmetadata.WithForwardedAuth(ctx, tf.c.AllowInsecureForCredentials())
	if err != nil {
		return nil, err
	}
	return ttnpb.NewDeviceRepositoryClient(conn).GetTemplate(ctx, in, opt)
}

type templateFetcherContextKeyType struct{}

var templateFetcherContextKey templateFetcherContextKeyType

// NewContextWithFetcher returns a context with the TemplateFetcher.
func NewContextWithFetcher(ctx context.Context, tf TemplateFetcher) context.Context {
	return context.WithValue(ctx, templateFetcherContextKey, tf)
}

// fetcherFromContext returns a TemplateFetcher from the context provided. Will return false if not present.
func fetcherFromContext(ctx context.Context) (TemplateFetcher, bool) {
	if fetcher, ok := ctx.Value(templateFetcherContextKey).(TemplateFetcher); ok {
		return fetcher, true
	}
	return nil, false
}
