package session

import (
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var evtUserLoginFailed = events.Define(
	"account.user.login_failed", "login user failure",
	events.WithVisibility(ttnpb.Right_RIGHT_USER_ALL),
	events.WithAuthFromContext(),
	events.WithClientInfoFromContext(),
)
