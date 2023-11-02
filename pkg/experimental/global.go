package experimental

var globalRegistry = NewRegistry()

// EnableFeatures enables the given features on the global registry.
func EnableFeatures(features ...string) {
	globalRegistry.EnableFeatures(features...)
}

// DisableFeatures disables the given features on the global registry.
func DisableFeatures(features ...string) {
	globalRegistry.DisableFeatures(features...)
}
