package gatewayconfigurationserver

// TheThingsGatewayConfig is the configuration for The Things Gateway.
type TheThingsGatewayConfig struct {
	Default struct {
		UpdateChannel string `name:"update-channel" description:"The default update channel that the gateways should use"`
		MQTTServer    string `name:"mqtt-server" description:"The default MQTT server that the gateways should use"`
		FirmwareURL   string `name:"firmware-url" description:"The default URL to the firmware storage"`
	} `name:"default" description:"Default gateway settings"`
}
