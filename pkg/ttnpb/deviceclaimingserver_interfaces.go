package ttnpb

// IDString implements the ttnpb.IDStringer interface.
func (req *AuthorizeApplicationRequest) IDString() string {
	appIDs := req.GetApplicationIds()
	if appIDs != nil {
		return appIDs.IDString()
	}
	return ""
}

// EntityType implements the ttnpb.IDStringer interface.
func (req *AuthorizeApplicationRequest) EntityType() string {
	appIDs := req.GetApplicationIds()
	if appIDs != nil {
		return appIDs.EntityType()
	}
	return ""
}

// IDString implements the ttnpb.IDStringer interface.
func (req *AuthorizeGatewayRequest) IDString() string {
	gtwIDs := req.GetGatewayIds()
	if gtwIDs != nil {
		return gtwIDs.IDString()
	}
	return ""
}

// EntityType implements the ttnpb.IDStringer interface.
func (req *AuthorizeGatewayRequest) EntityType() string {
	gtwIDs := req.GetGatewayIds()
	if gtwIDs != nil {
		return gtwIDs.EntityType()
	}
	return ""
}
