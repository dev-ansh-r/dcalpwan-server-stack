package auth

import (
	"context"
	"crypto/x509/pkix"
)

type x509DNKeyType struct{}

var x509DNKey x509DNKeyType

// NewContextWithX509DN returns a new context with the given distinguished name.
func NewContextWithX509DN(ctx context.Context, name pkix.Name) context.Context {
	return context.WithValue(ctx, x509DNKey, name)
}

// X509DNFromContext returns the distinguished name from the given context.
func X509DNFromContext(ctx context.Context) (pkix.Name, bool) {
	if name, ok := ctx.Value(x509DNKey).(pkix.Name); ok {
		return name, true
	}
	return pkix.Name{}, false
}
