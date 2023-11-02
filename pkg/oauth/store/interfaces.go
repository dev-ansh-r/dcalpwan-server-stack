package store

import (
	"context"

	"go.thethings.network/lorawan-stack/v3/pkg/identityserver/store"
)

// Interface is the store used by the account app.
type Interface interface {
	store.UserStore
	store.UserSessionStore

	store.ClientStore
	store.OAuthStore
}

// TransactionalInterface is Interface, but with a method that uses a transaction.
type TransactionalInterface interface {
	Interface

	// Transact runs a transaction using the store.
	Transact(context.Context, func(context.Context, Interface) error) error
}
