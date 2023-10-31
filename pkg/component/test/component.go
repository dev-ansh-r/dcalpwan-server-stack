package test

import (
	"testing"

	"go.thethings.network/lorawan-stack/v3/pkg/component"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
)

// NewComponent returns a new Component that can be used for testing.
func NewComponent(tb testing.TB, config *component.Config, opts ...component.Option) *component.Component {
	c, err := component.New(test.GetLogger(tb), config, opts...)
	if err != nil {
		tb.Fatalf("Failed to create component: %v", err)
	}
	return c
}

// StartComponent starts the component for testing.
func StartComponent(tb testing.TB, c *component.Component) {
	if err := c.Start(); err != nil {
		tb.Fatalf("Failed to start component: %v", err)
	}
}
