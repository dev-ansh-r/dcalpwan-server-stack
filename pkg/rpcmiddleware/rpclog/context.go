package rpclog

import "context"

type requestFieldsKeyType struct{}

var requestFieldsKey requestFieldsKeyType

type requestFieldsValue struct {
	fields map[string]any
}

func requestFieldsFromContext(ctx context.Context) (*requestFieldsValue, bool) {
	value, ok := ctx.Value(requestFieldsKey).(*requestFieldsValue)
	return value, ok
}

func newContextWithRequestFields(parent context.Context) context.Context {
	return context.WithValue(parent, requestFieldsKey, &requestFieldsValue{
		fields: make(map[string]any),
	})
}

// AddField adds a log field to the fields in the request context.
// Not safe for concurrent use.
func AddField(ctx context.Context, key string, value any) {
	if v, ok := requestFieldsFromContext(ctx); ok {
		v.fields[key] = value
	}
}
