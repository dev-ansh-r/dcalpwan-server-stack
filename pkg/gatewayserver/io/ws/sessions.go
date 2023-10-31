package ws

import (
	"context"
	"sync"
)

type sessionKeyType struct{}

var sessionKey sessionKeyType

// Session contains the session state for a single gateway.
type Session struct {
	DataMu sync.RWMutex
	Data   any
}

// NewContextWithSession returns a new context with the session.
func NewContextWithSession(ctx context.Context, session *Session) context.Context {
	return context.WithValue(ctx, sessionKey, session)
}

// SessionFromContext returns a new session from the context.
// The session value can be modified by the caller.
func SessionFromContext(ctx context.Context) *Session {
	if session, ok := ctx.Value(sessionKey).(*Session); ok {
		return session
	}
	return nil
}
