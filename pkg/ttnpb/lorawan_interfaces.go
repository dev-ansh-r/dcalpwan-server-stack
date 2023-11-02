package ttnpb

// All EntityType methods implement the IDStringer interface.

func (m *GatewayAntennaIdentifiers) EntityType() string {
	return m.GetGatewayIds().EntityType()
}

func (m *UplinkToken) EntityType() string {
	return m.GetIds().EntityType()
}

// All IDString methods implement the IDStringer interface.

func (m *GatewayAntennaIdentifiers) IDString() string {
	return m.GetGatewayIds().IDString()
}

func (m *UplinkToken) IDString() string {
	return m.GetIds().IDString()
}

// All ExtractRequestFields methods are used by github.com/grpc-ecosystem/go-grpc-middleware/tags.

func (m *GatewayAntennaIdentifiers) ExtractRequestFields(dst map[string]interface{}) {
	m.GetGatewayIds().ExtractRequestFields(dst)
}

func (m *UplinkToken) ExtractRequestFields(dst map[string]interface{}) {
	m.GetIds().ExtractRequestFields(dst)
}
