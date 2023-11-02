package ttnpb

import (
	"context"
)

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *CreateOrganizationRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetOrganizationRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (m *UpdateOrganizationRequest) ValidateContext(context.Context) error {
	if len(m.FieldMask.GetPaths()) == 0 {
		return m.ValidateFields()
	}
	return m.ValidateFields(append(FieldsWithPrefix("organization", m.FieldMask.GetPaths()...),
		"organization.ids",
	)...)
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *CreateOrganizationAPIKeyRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *ListOrganizationAPIKeysRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetOrganizationAPIKeyRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *UpdateOrganizationAPIKeyRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetOrganizationCollaboratorRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *SetOrganizationCollaboratorRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *ListOrganizationCollaboratorsRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}
