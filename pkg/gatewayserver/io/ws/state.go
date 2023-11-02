package ws

import (
	"context"
)

// state represents the LBS session state.
type state struct {
	ID       *int32
	TimeSync *bool
}

// updateState updates the session state.
func updateState(ctx context.Context, f func(*state)) {
	session := SessionFromContext(ctx)
	session.DataMu.Lock()
	defer session.DataMu.Unlock()
	st, ok := session.Data.(*state)
	if !ok {
		st = &state{}
		session.Data = st
	}
	f(st)
}

// GetState returns the session state.
func getState(ctx context.Context, f func(*state) any) any {
	session := SessionFromContext(ctx)
	session.DataMu.RLock()
	defer session.DataMu.RUnlock()
	st, ok := session.Data.(*state)
	if !ok {
		return nil
	}
	return f(st)
}

// UpdateSessionID updates the session ID.
func UpdateSessionID(ctx context.Context, id int32) {
	updateState(ctx, func(st *state) {
		st.ID = &id
	})
}

// UpdateSessionTimeSync updates the session time sync.
func UpdateSessionTimeSync(ctx context.Context, b bool) {
	updateState(ctx, func(st *state) {
		st.TimeSync = &b
	})
}

// GetSessionID returns the session ID.
func GetSessionID(ctx context.Context) (int32, bool) {
	i, ok := getState(ctx, func(st *state) any {
		if st.ID != nil {
			return *st.ID
		}
		return nil
	}).(int32)
	return i, ok
}

// GetSessionTimeSync returns the session time sync.
func GetSessionTimeSync(ctx context.Context) (enabled bool, ok bool) {
	d, ok := getState(ctx, func(st *state) any {
		if st.TimeSync != nil {
			return *st.TimeSync
		}
		return nil
	}).(bool)
	return d, ok
}
