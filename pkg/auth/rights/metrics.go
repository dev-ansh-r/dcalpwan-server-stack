package rights

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/metrics"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

const subsystem = "rights_hook"

var rightsRequests = metrics.NewContextualCounterVec(
	prometheus.CounterOpts{
		Subsystem: subsystem,
		Name:      "requests_total",
		Help:      "Rights Requests",
	},
	[]string{"type", "result"},
)

var rightsFetches = metrics.NewContextualCounterVec(
	prometheus.CounterOpts{
		Subsystem: subsystem,
		Name:      "fetches_total",
		Help:      "Rights Fetches from Identity Server",
	},
	[]string{"type", "result"},
)

func init() {
	metrics.MustRegister(rightsRequests, rightsFetches)
}

func registerEntityRights(ctx context.Context, c *metrics.ContextualCounterVec, entity string, rights *ttnpb.Rights, err error) {
	switch {
	case errors.IsUnauthenticated(err):
		c.WithLabelValues(ctx, entity, "unauthenticated").Inc()
	case errors.IsPermissionDenied(err):
		c.WithLabelValues(ctx, entity, "permission_denied").Inc()
	case err != nil:
		c.WithLabelValues(ctx, entity, "error").Inc()
	case rights == nil || len(rights.Rights) == 0:
		c.WithLabelValues(ctx, entity, "zero").Inc()
	default:
		c.WithLabelValues(ctx, entity, "ok").Inc()
	}
}

func registerRightsRequest(ctx context.Context, entity string, rights *ttnpb.Rights, err error) {
	registerEntityRights(ctx, rightsRequests, entity, rights, err)
}

func registerRightsFetch(ctx context.Context, entity string, rights *ttnpb.Rights, err error) {
	registerEntityRights(ctx, rightsFetches, entity, rights, err)
}

func registerAuthInfo(ctx context.Context, c *metrics.ContextualCounterVec, authInfo *ttnpb.AuthInfoResponse, err error) {
	switch {
	case errors.IsUnauthenticated(err):
		c.WithLabelValues(ctx, "auth_info", "unauthenticated").Inc()
	case errors.IsPermissionDenied(err):
		c.WithLabelValues(ctx, "auth_info", "permission_denied").Inc()
	case err != nil:
		c.WithLabelValues(ctx, "auth_info", "error").Inc()
	case authInfo.GetIsAdmin():
		c.WithLabelValues(ctx, "auth_info", "admin").Inc()
	case authInfo.GetUniversalRights() == nil || len(authInfo.GetUniversalRights().GetRights()) == 0:
		c.WithLabelValues(ctx, "auth_info", "zero").Inc()
	default:
		c.WithLabelValues(ctx, "auth_info", "ok").Inc()
	}
}

func registerAuthInfoRequest(ctx context.Context, authInfo *ttnpb.AuthInfoResponse, err error) {
	registerAuthInfo(ctx, rightsRequests, authInfo, err)
}

func registerAuthInfoFetch(ctx context.Context, authInfo *ttnpb.AuthInfoResponse, err error) {
	registerAuthInfo(ctx, rightsFetches, authInfo, err)
}
