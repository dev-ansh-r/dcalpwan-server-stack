package ttnpb

import "context"

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *CreateClientRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetClientRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (m *UpdateClientRequest) ValidateContext(context.Context) error {
	if len(m.FieldMask.GetPaths()) == 0 {
		return m.ValidateFields()
	}
	return m.ValidateFields(append(FieldsWithPrefix("client", m.FieldMask.GetPaths()...),
		"client.ids",
	)...)
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetClientCollaboratorRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *SetClientCollaboratorRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *ListClientCollaboratorsRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}
