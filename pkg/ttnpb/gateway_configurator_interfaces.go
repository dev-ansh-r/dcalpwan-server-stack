package ttnpb

// Implement the IDStringer interface.

// EntityType implements IDStringer.
func (req *PullGatewayConfigurationRequest) EntityType() string {
	gtwIDs := req.GetGatewayIds()
	if gtwIDs == nil {
		return ""
	}
	return gtwIDs.EntityType()
}

// IDString implements IDStringer.
func (req *PullGatewayConfigurationRequest) IDString() string {
	gtwIDs := req.GetGatewayIds()
	if gtwIDs == nil {
		return ""
	}
	return gtwIDs.IDString()
}
