package packetbrokeragent

import "context"

type contextDecoupler interface {
	FromRequestContext(context.Context) context.Context
}
