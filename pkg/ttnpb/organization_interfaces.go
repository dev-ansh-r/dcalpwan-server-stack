package ttnpb

// All EntityType methods implement the IDStringer interface.

func (m *Organization) EntityType() string {
	return m.GetIds().EntityType()
}

func (m *GetOrganizationRequest) EntityType() string {
	return m.GetOrganizationIds().EntityType()
}

func (m *CreateOrganizationRequest) EntityType() string {
	return m.GetOrganization().GetIds().EntityType()
}

func (m *UpdateOrganizationRequest) EntityType() string {
	return m.GetOrganization().GetIds().EntityType()
}

func (m *ListOrganizationAPIKeysRequest) EntityType() string {
	return m.GetOrganizationIds().EntityType()
}

func (m *GetOrganizationAPIKeyRequest) EntityType() string {
	return m.GetOrganizationIds().EntityType()
}

func (m *CreateOrganizationAPIKeyRequest) EntityType() string {
	return m.GetOrganizationIds().EntityType()
}

func (m *UpdateOrganizationAPIKeyRequest) EntityType() string {
	return m.GetOrganizationIds().EntityType()
}

func (m *ListOrganizationCollaboratorsRequest) EntityType() string {
	return m.GetOrganizationIds().EntityType()
}

func (m *GetOrganizationCollaboratorRequest) EntityType() string {
	return m.GetOrganizationIds().EntityType()
}

func (m *SetOrganizationCollaboratorRequest) EntityType() string {
	return m.GetOrganizationIds().EntityType()
}

// All IDString methods implement the IDStringer interface.

func (m *Organization) IDString() string {
	return m.GetIds().IDString()
}

func (m *GetOrganizationRequest) IDString() string {
	return m.GetOrganizationIds().IDString()
}

func (m *CreateOrganizationRequest) IDString() string {
	return m.GetOrganization().GetIds().IDString()
}

func (m *UpdateOrganizationRequest) IDString() string {
	return m.GetOrganization().GetIds().IDString()
}

func (m *ListOrganizationAPIKeysRequest) IDString() string {
	return m.GetOrganizationIds().IDString()
}

func (m *GetOrganizationAPIKeyRequest) IDString() string {
	return m.GetOrganizationIds().IDString()
}

func (m *CreateOrganizationAPIKeyRequest) IDString() string {
	return m.GetOrganizationIds().IDString()
}

func (m *UpdateOrganizationAPIKeyRequest) IDString() string {
	return m.GetOrganizationIds().IDString()
}

func (m *ListOrganizationCollaboratorsRequest) IDString() string {
	return m.GetOrganizationIds().IDString()
}

func (m *GetOrganizationCollaboratorRequest) IDString() string {
	return m.GetOrganizationIds().IDString()
}

func (m *SetOrganizationCollaboratorRequest) IDString() string {
	return m.GetOrganizationIds().IDString()
}

// All ExtractRequestFields methods are used by github.com/grpc-ecosystem/go-grpc-middleware/tags.

func (m *Organization) ExtractRequestFields(dst map[string]interface{}) {
	m.GetIds().ExtractRequestFields(dst)
}

func (m *GetOrganizationRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetOrganizationIds().ExtractRequestFields(dst)
}

func (m *UpdateOrganizationRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetOrganization().GetIds().ExtractRequestFields(dst)
}

func (m *ListOrganizationAPIKeysRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetOrganizationIds().ExtractRequestFields(dst)
}

func (m *GetOrganizationAPIKeyRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetOrganizationIds().ExtractRequestFields(dst)
}

func (m *CreateOrganizationAPIKeyRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetOrganizationIds().ExtractRequestFields(dst)
}

func (m *UpdateOrganizationAPIKeyRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetOrganizationIds().ExtractRequestFields(dst)
}

func (m *ListOrganizationCollaboratorsRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetOrganizationIds().ExtractRequestFields(dst)
}

func (m *GetOrganizationCollaboratorRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetOrganizationIds().ExtractRequestFields(dst)
}

// Wrap methods of m.OrganizationIdentifiers.

func (m *Organization) GetEntityIdentifiers() *EntityIdentifiers {
	return m.GetIds().GetEntityIdentifiers()
}

func (m *Organization) GetOrganizationOrUserIdentifiers() *OrganizationOrUserIdentifiers {
	return m.GetIds().GetOrganizationOrUserIdentifiers()
}
