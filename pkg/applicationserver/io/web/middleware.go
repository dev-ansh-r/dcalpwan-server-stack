package web

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.thethings.network/lorawan-stack/v3/pkg/auth"
	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/ratelimit"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
)

type (
	deviceIDKeyType  struct{}
	webhookIDKeyType struct{}
)

var (
	deviceIDKey  deviceIDKeyType
	webhookIDKey webhookIDKeyType
)

func withDeviceID(ctx context.Context, id *ttnpb.EndDeviceIdentifiers) context.Context {
	return context.WithValue(ctx, deviceIDKey, id)
}

func deviceIDFromContext(ctx context.Context) *ttnpb.EndDeviceIdentifiers {
	id, ok := ctx.Value(deviceIDKey).(*ttnpb.EndDeviceIdentifiers)
	if !ok {
		panic("no end device identifiers found in context")
	}
	return id
}

func withWebhookID(ctx context.Context, id *ttnpb.ApplicationWebhookIdentifiers) context.Context {
	return context.WithValue(ctx, webhookIDKey, id)
}

func webhookIDFromContext(ctx context.Context) *ttnpb.ApplicationWebhookIdentifiers {
	id, ok := ctx.Value(webhookIDKey).(*ttnpb.ApplicationWebhookIdentifiers)
	if !ok {
		panic("no webhook identifiers found in context")
	}
	return id
}

func (w *webhooks) validateAndFillIDs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		vars := mux.Vars(r)
		appID := ttnpb.ApplicationIdentifiers{
			ApplicationId: vars["application_id"],
		}

		devID := &ttnpb.EndDeviceIdentifiers{
			ApplicationIds: &appID,
			DeviceId:       vars["device_id"],
		}
		if err := devID.ValidateContext(ctx); err != nil {
			webhandlers.Error(w, r, err)
			return
		}
		ctx = withDeviceID(ctx, devID)

		hookID := &ttnpb.ApplicationWebhookIdentifiers{
			ApplicationIds: &appID,
			WebhookId:      vars["webhook_id"],
		}
		if err := hookID.ValidateContext(ctx); err != nil {
			webhandlers.Error(w, r, err)
			return
		}
		ctx = withWebhookID(ctx, hookID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (w *webhooks) requireApplicationRights(required ...ttnpb.Right) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			appID := deviceIDFromContext(ctx).ApplicationIds
			if err := rights.RequireApplication(ctx, appID, required...); err != nil {
				webhandlers.Error(res, req, err)
				return
			}
			next.ServeHTTP(res, req)
		})
	}
}

var errRateLimitExceeded = errors.DefineResourceExhausted("rate_limit_exceeded", "rate limit exceeded")

func (w *webhooks) requireRateLimits() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			authTokenID := ""
			token := req.Header.Get("Authorization")
			if _, v, _, err := auth.SplitToken(token); err == nil && v != "" {
				authTokenID = v
			}

			resource := ratelimit.ApplicationWebhooksDownResource(ctx, deviceIDFromContext(ctx), authTokenID)
			limit, result := w.server.RateLimiter().RateLimit(resource)
			result.SetHTTPHeaders(res.Header())
			if limit {
				webhandlers.Error(res, req, errRateLimitExceeded.New())
				return
			}

			next.ServeHTTP(res, req)
		})
	}
}
