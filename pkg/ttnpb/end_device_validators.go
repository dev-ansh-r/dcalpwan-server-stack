package ttnpb

import "context"

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *CreateEndDeviceRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *ListEndDevicesRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (m *UpdateEndDeviceRequest) ValidateContext(context.Context) error {
	if len(m.FieldMask.GetPaths()) == 0 {
		return m.ValidateFields()
	}
	return m.ValidateFields(append(FieldsWithPrefix("end_device", m.FieldMask.GetPaths()...),
		"end_device.ids.application_ids",
		"end_device.ids.device_id",
	)...)
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (m *SetEndDeviceRequest) ValidateContext(context.Context) error {
	if len(m.FieldMask.GetPaths()) == 0 {
		return m.ValidateFields()
	}
	return m.ValidateFields(append(FieldsWithPrefix("end_device", m.FieldMask.GetPaths()...),
		"end_device.ids.application_ids",
		"end_device.ids.device_id",
	)...)
}

// ValidateContext wraps the generated validator with (optionally context-based) custom checks.
func (req *ResetAndGetEndDeviceRequest) ValidateContext(context.Context) error {
	return req.ValidateFields()
}
