package ttnpb

import "context"

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *CreateGatewayRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetGatewayRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (m *UpdateGatewayRequest) ValidateContext(context.Context) error {
	if len(m.FieldMask.GetPaths()) == 0 {
		return m.ValidateFields()
	}
	return m.ValidateFields(append(FieldsWithPrefix("gateway", m.FieldMask.GetPaths()...),
		"gateway.ids",
	)...)
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *CreateGatewayAPIKeyRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *ListGatewayAPIKeysRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetGatewayAPIKeyRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *UpdateGatewayAPIKeyRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetGatewayCollaboratorRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *SetGatewayCollaboratorRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *ListGatewayCollaboratorsRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *PullGatewayConfigurationRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}
