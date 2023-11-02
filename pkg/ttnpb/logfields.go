package ttnpb

func (ids *ApplicationIdentifiers) ExtractRequestFields(m map[string]interface{}) {
	if ids == nil {
		return
	}
	m["application_id"] = ids.ApplicationId
}

func (ids *ClientIdentifiers) ExtractRequestFields(m map[string]interface{}) {
	if ids == nil {
		return
	}
	m["client_id"] = ids.ClientId
}

func (ids *EndDeviceIdentifiers) ExtractRequestFields(m map[string]interface{}) {
	if ids == nil {
		return
	}
	m["application_id"] = ids.GetApplicationIds().GetApplicationId()
	m["device_id"] = ids.DeviceId
}

func (ids *GatewayIdentifiers) ExtractRequestFields(m map[string]interface{}) {
	if ids == nil {
		return
	}
	m["gateway_id"] = ids.GatewayId
}

func (ids *OrganizationIdentifiers) ExtractRequestFields(m map[string]interface{}) {
	if ids == nil {
		return
	}
	m["organization_id"] = ids.OrganizationId
}

func (ids *UserIdentifiers) ExtractRequestFields(m map[string]interface{}) {
	if ids == nil {
		return
	}
	m["user_id"] = ids.UserId
}

func extractCollaboratorFields(m map[string]interface{}, ids *OrganizationOrUserIdentifiers) {
	if ids == nil {
		return
	}
	switch oneof := ids.Ids.(type) {
	case nil:
	case *OrganizationOrUserIdentifiers_OrganizationIds:
		m["collaborator_organization_id"] = oneof.OrganizationIds.OrganizationId
	case *OrganizationOrUserIdentifiers_UserIds:
		m["collaborator_user_id"] = oneof.UserIds.UserId
	default:
		panic("missed oneof type in extractCollaboratorFields()")
	}
}

func (req *CreateApplicationRequest) ExtractRequestFields(m map[string]interface{}) {
	if req == nil {
		return
	}
	req.Application.ExtractRequestFields(m)
	extractCollaboratorFields(m, req.GetCollaborator())
}

func (req *CreateClientRequest) ExtractRequestFields(m map[string]interface{}) {
	if req == nil {
		return
	}
	req.Client.ExtractRequestFields(m)
	extractCollaboratorFields(m, req.GetCollaborator())
}

func (req *CreateGatewayRequest) ExtractRequestFields(m map[string]interface{}) {
	if req == nil {
		return
	}
	req.Gateway.ExtractRequestFields(m)
	extractCollaboratorFields(m, req.GetCollaborator())
}

func (req *CreateOrganizationRequest) ExtractRequestFields(m map[string]interface{}) {
	if req == nil {
		return
	}
	req.Organization.ExtractRequestFields(m)
	extractCollaboratorFields(m, req.GetCollaborator())
}

func (req *SetApplicationCollaboratorRequest) ExtractRequestFields(m map[string]interface{}) {
	if req == nil {
		return
	}
	req.ApplicationIds.ExtractRequestFields(m)
	extractCollaboratorFields(m, req.GetCollaborator().GetIds())
}

func (req *SetClientCollaboratorRequest) ExtractRequestFields(m map[string]interface{}) {
	if req == nil {
		return
	}
	req.GetClientIds().ExtractRequestFields(m)
	extractCollaboratorFields(m, req.GetCollaborator().GetIds())
}

func (req *SetGatewayCollaboratorRequest) ExtractRequestFields(m map[string]interface{}) {
	if req == nil {
		return
	}
	req.GetGatewayIds().ExtractRequestFields(m)
	extractCollaboratorFields(m, req.GetCollaborator().GetIds())
}

func (req *SetOrganizationCollaboratorRequest) ExtractRequestFields(m map[string]interface{}) {
	if req == nil {
		return
	}
	req.GetOrganizationIds().ExtractRequestFields(m)
	extractCollaboratorFields(m, req.GetCollaborator().GetIds())
}
