package events

import "context"

// ContextMarshaler interface for marshaling/unmarshaling contextual information
// to/from events.
type ContextMarshaler interface {
	MarshalContext(context.Context) []byte
	UnmarshalContext(context.Context, []byte) (context.Context, error)
}

var contextMarshalers = map[string]ContextMarshaler{}

// RegisterContextMarshaler registers a ContextMarshaler with the given name.
// This should only be called from init funcs.
func RegisterContextMarshaler(name string, m ContextMarshaler) {
	contextMarshalers[name] = m
}

func unmarshalContext(ctx context.Context, data map[string][]byte) (context.Context, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	var err error
	for name, payload := range data {
		m, ok := contextMarshalers[name]
		if !ok {
			continue
		}
		ctx, err = m.UnmarshalContext(ctx, payload)
		if err != nil {
			return nil, err
		}
	}
	return ctx, nil
}

func marshalContext(ctx context.Context) (map[string][]byte, error) {
	data := make(map[string][]byte, len(contextMarshalers))
	for name, m := range contextMarshalers {
		payload := m.MarshalContext(ctx)
		if payload == nil {
			continue
		}
		data[name] = payload
	}
	return data, nil
}
