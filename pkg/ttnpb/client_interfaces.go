package ttnpb

// All EntityType methods implement the IDStringer interface.

func (m *Client) EntityType() string {
	return m.GetIds().EntityType()
}

func (m *GetClientRequest) EntityType() string {
	return m.GetClientIds().EntityType()
}

func (m *CreateClientRequest) EntityType() string {
	return m.GetClient().EntityType()
}

func (m *UpdateClientRequest) EntityType() string {
	return m.GetClient().EntityType()
}

func (m *ListClientCollaboratorsRequest) EntityType() string {
	return m.GetClientIds().EntityType()
}

func (m *GetClientCollaboratorRequest) EntityType() string {
	return m.GetClientIds().EntityType()
}

func (m *SetClientCollaboratorRequest) EntityType() string {
	return m.GetClientIds().EntityType()
}

// All IDString methods implement the IDStringer interface.

func (m *Client) IDString() string {
	return m.GetIds().IDString()
}

func (m *GetClientRequest) IDString() string {
	return m.GetClientIds().IDString()
}

func (m *CreateClientRequest) IDString() string {
	return m.GetClient().IDString()
}

func (m *UpdateClientRequest) IDString() string {
	return m.GetClient().IDString()
}

func (m *ListClientCollaboratorsRequest) IDString() string {
	return m.GetClientIds().IDString()
}

func (m *GetClientCollaboratorRequest) IDString() string {
	return m.GetClientIds().IDString()
}

func (m *SetClientCollaboratorRequest) IDString() string {
	return m.GetClientIds().IDString()
}

// All ExtractRequestFields methods are used by github.com/grpc-ecosystem/go-grpc-middleware/tags.

func (m *Client) ExtractRequestFields(dst map[string]interface{}) {
	m.GetIds().ExtractRequestFields(dst)
}

func (m *GetClientRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetClientIds().ExtractRequestFields(dst)
}

func (m *UpdateClientRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetClient().ExtractRequestFields(dst)
}

func (m *ListClientCollaboratorsRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetClientIds().ExtractRequestFields(dst)
}

func (m *GetClientCollaboratorRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetClientIds().ExtractRequestFields(dst)
}

// Wrap methods of m.ClientIdentifiers.

func (m *Client) GetEntityIdentifiers() *EntityIdentifiers {
	if m == nil {
		return nil
	}
	return m.GetIds().GetEntityIdentifiers()
}
