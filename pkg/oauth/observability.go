package oauth

import (
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	// EvtUserLogin indicates a user login.
	EvtUserLogin = events.Define(
		"oauth.user.login", "login user successful",
		events.WithVisibility(ttnpb.Right_RIGHT_USER_ALL),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
	)
	// EvtUserLogout indicates a user logout.
	EvtUserLogout = events.Define(
		"oauth.user.logout", "logout user",
		events.WithVisibility(ttnpb.Right_RIGHT_USER_ALL),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
	)

	evtAuthorize = events.Define(
		"oauth.authorize", "authorize OAuth client",
		events.WithVisibility(ttnpb.Right_RIGHT_USER_AUTHORIZED_CLIENTS),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
	)
	evtTokenExchange = events.Define(
		"oauth.token.exchange", "exchange OAuth access token",
		events.WithVisibility(ttnpb.Right_RIGHT_USER_AUTHORIZED_CLIENTS),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
	)
	evtAccessTokenDeleted = events.Define(
		"oauth.token.deleted", "delete access token",
		events.WithVisibility(ttnpb.Right_RIGHT_USER_AUTHORIZED_CLIENTS),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
	)
	evtUserSessionTerminated = events.Define(
		"oauth.session.terminated", "terminate user session",
		events.WithVisibility(ttnpb.Right_RIGHT_USER_AUTHORIZED_CLIENTS),
		events.WithAuthFromContext(),
		events.WithClientInfoFromContext(),
	)
)
