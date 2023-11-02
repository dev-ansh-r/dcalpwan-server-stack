package ttnpb

import "context"

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *CryptoServicePayloadRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *JoinAcceptMICRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *DeriveSessionKeysRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetRootKeysRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *ProvisionEndDevicesRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *GetApplicationActivationSettingsRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *SetApplicationActivationSettingsRequest) ValidateContext(context.Context) error {
	if len(req.FieldMask.GetPaths()) == 0 {
		return req.ValidateFields()
	}
	return req.ValidateFields(append(FieldsWithPrefix("settings", req.FieldMask.GetPaths()...),
		"application_ids",
	)...)
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *DeleteApplicationActivationSettingsRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}
