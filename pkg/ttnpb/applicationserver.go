package ttnpb

// All EntityType methods implement the IDStringer interface.

func (m *GetApplicationLinkRequest) EntityType() string {
	return m.GetApplicationIds().EntityType()
}

func (m *SetApplicationLinkRequest) EntityType() string {
	return m.GetApplicationIds().EntityType()
}

// All IDString methods implement the IDStringer interface.

func (m *GetApplicationLinkRequest) IDString() string {
	return m.GetApplicationIds().IDString()
}

func (m *SetApplicationLinkRequest) IDString() string {
	return m.GetApplicationIds().IDString()
}

// All ExtractRequestFields methods are used by github.com/grpc-ecosystem/go-grpc-middleware/tags.

func (m *GetApplicationLinkRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetApplicationIds().ExtractRequestFields(dst)
}

func (m *SetApplicationLinkRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetApplicationIds().ExtractRequestFields(dst)
}
