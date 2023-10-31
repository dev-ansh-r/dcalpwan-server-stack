package blocklist

import "context"

type ctxKeyType struct{}

var ctxKey ctxKeyType

// FromContext returns the blocklists from the given context.
func FromContext(ctx context.Context) Blocklists {
	if blocklists, ok := ctx.Value(ctxKey).(Blocklists); ok {
		return blocklists
	}
	return nil
}

// NewContext returns a new context derived from parent with the given blocklists attached.
func NewContext(parent context.Context, blocklists ...*Blocklist) context.Context {
	if len(blocklists) == 0 {
		return parent
	}
	parentBlocklists := FromContext(parent)
	newBlocklists := make(Blocklists, 0, len(parentBlocklists)+len(blocklists))
	newBlocklists = append(newBlocklists, parentBlocklists...)
	newBlocklists = append(newBlocklists, blocklists...)
	return context.WithValue(parent, ctxKey, newBlocklists)
}
