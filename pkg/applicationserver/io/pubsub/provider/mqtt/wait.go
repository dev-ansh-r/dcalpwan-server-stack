package mqtt

import (
	"context"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// waitToken awaits the token operation to finish in parallel with the context.
// Keep in mind that context deadlines do not cancel the underlying MQTT client operation.
func waitToken(ctx context.Context, token mqtt.Token) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-token.Done():
		return token.Error()
	}
}
