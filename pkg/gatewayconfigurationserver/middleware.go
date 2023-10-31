package gatewayconfigurationserver

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"go.thethings.network/lorawan-stack/v3/pkg/auth/rights"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/webhandlers"
	"go.thethings.network/lorawan-stack/v3/pkg/webmiddleware"
)

type gatewayIDKeyType struct{}

var gatewayIDKey gatewayIDKeyType

func withGatewayID(ctx context.Context, id *ttnpb.GatewayIdentifiers) context.Context {
	return context.WithValue(ctx, gatewayIDKey, id)
}

func gatewayIDFromContext(ctx context.Context) *ttnpb.GatewayIdentifiers {
	id, ok := ctx.Value(gatewayIDKey).(*ttnpb.GatewayIdentifiers)
	if !ok {
		panic("no gateway identifiers found in context")
	}
	return id
}

func validateAndFillIDs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		gtwID := &ttnpb.GatewayIdentifiers{
			GatewayId: mux.Vars(r)["gateway_id"],
		}
		if err := gtwID.ValidateContext(ctx); err != nil {
			webhandlers.Error(w, r, err)
			return
		}
		ctx = withGatewayID(ctx, gtwID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) requireGatewayRights(required ...ttnpb.Right) webmiddleware.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			gtwID := gatewayIDFromContext(ctx)
			if err := rights.RequireGateway(ctx, gtwID, required...); err != nil {
				webhandlers.Error(w, r, err)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
