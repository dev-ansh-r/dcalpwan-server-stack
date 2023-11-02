package ttnpb

// GetEntityIdentifiers returns the CollaboratorOf field as EntityIdentifiers.
func (req *SearchAccountsRequest) GetEntityIdentifiers() *EntityIdentifiers {
	if req == nil || req.CollaboratorOf == nil {
		return nil
	}
	switch v := req.CollaboratorOf.(type) {
	default:
		return nil
	case *SearchAccountsRequest_ApplicationIds:
		return v.ApplicationIds.GetEntityIdentifiers()
	case *SearchAccountsRequest_ClientIds:
		return v.ClientIds.GetEntityIdentifiers()
	case *SearchAccountsRequest_GatewayIds:
		return v.GatewayIds.GetEntityIdentifiers()
	case *SearchAccountsRequest_OrganizationIds:
		return v.OrganizationIds.GetEntityIdentifiers()
	}
}
