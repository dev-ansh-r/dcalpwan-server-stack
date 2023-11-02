package store

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
)

// Interface is the store used by the account app.
type Interface interface {
	// UserStore, LoginTokenStore and UserSessionStore are needed for user login/logout.
	store.UserStore
	store.LoginTokenStore
	store.UserSessionStore
}

// TransactionalStore is Interface, but with a method that uses a transaction.
type TransactionalInterface interface {
	Interface

	// Transact runs a transaction using the store.
	Transact(context.Context, func(context.Context, Interface) error) error
}

// WithSoftDeleted returns a context that tells the store to include (only) deleted entities.
var WithSoftDeleted = store.WithSoftDeleted
