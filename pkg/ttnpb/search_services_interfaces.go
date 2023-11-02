package ttnpb

// EntityType implements the IDStringer interface.
func (m *SearchEndDevicesRequest) EntityType() string {
	return m.GetApplicationIds().EntityType()
}

// IDString implements the IDStringer interface.
func (m *SearchEndDevicesRequest) IDString() string {
	return m.GetApplicationIds().IDString()
}

// ExtractRequestFields is used by github.com/grpc-ecosystem/go-grpc-middleware/tags.
func (m *SearchEndDevicesRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetApplicationIds().ExtractRequestFields(dst)
}
