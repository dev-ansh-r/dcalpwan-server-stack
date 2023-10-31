package joinserver

import "go.thethings.network/lorawan-stack/v3/pkg/types"

// Config represents the JoinServer configuration.
type Config struct {
	Devices                       DeviceRegistry                       `name:"-"`
	Keys                          KeyRegistry                          `name:"-"`
	ApplicationActivationSettings ApplicationActivationSettingRegistry `name:"-"`
	JoinEUIPrefixes               []types.EUI64Prefix                  `name:"join-eui-prefix" description:"JoinEUI prefixes handled by this Join Server"`
	DefaultJoinEUI                types.EUI64                          `name:"default-join-eui" description:"Default JoinEUI for this Join Server"`
	DeviceKEKLabel                string                               `name:"device-kek-label" description:"Label of KEK used to encrypt device keys at rest"`
	DevNonceLimit                 int                                  `name:"dev-nonce-limit" description:"Amount of DevNonces stored per device"`
	SessionKeyLimit               int                                  `name:"session-key-limit" description:"Amount of session keys stored per device"`
}
