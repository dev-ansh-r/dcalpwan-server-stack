package ttnpb

// GetEntityIdentifiers returns the EntityIdentifiers for the used access method.
func (m *AuthInfoResponse) GetEntityIdentifiers() *EntityIdentifiers {
	if m == nil {
		return nil
	}
	switch accessMethod := m.GetAccessMethod().(type) {
	case *AuthInfoResponse_ApiKey:
		return accessMethod.ApiKey.EntityIds
	case *AuthInfoResponse_OauthAccessToken:
		return accessMethod.OauthAccessToken.UserIds.GetEntityIdentifiers()
	case *AuthInfoResponse_UserSession:
		return accessMethod.UserSession.GetUserIds().GetEntityIdentifiers()
	case *AuthInfoResponse_GatewayToken_:
		return accessMethod.GatewayToken.GetGatewayIds().GetEntityIdentifiers()
	}
	return nil
}

// GetRights returns the entity Rights for the used access method.
func (m *AuthInfoResponse) GetRights() []Right {
	if m == nil {
		return nil
	}
	switch accessMethod := m.GetAccessMethod().(type) {
	case *AuthInfoResponse_ApiKey:
		return accessMethod.ApiKey.GetApiKey().GetRights()
	case *AuthInfoResponse_OauthAccessToken:
		return accessMethod.OauthAccessToken.GetRights()
	case *AuthInfoResponse_UserSession:
		return RightsFrom(Right_RIGHT_ALL).Implied().GetRights()
	case *AuthInfoResponse_GatewayToken_:
		return accessMethod.GatewayToken.GetRights()
	}
	return nil
}

// GetOrganizationOrUserIdentifiers returns the OrganizationOrUserIdentifiers for the used access method.
func (m *AuthInfoResponse) GetOrganizationOrUserIdentifiers() *OrganizationOrUserIdentifiers {
	ids := m.GetEntityIdentifiers()
	if ids == nil {
		return nil
	}
	if ids := ids.GetOrganizationIds(); ids != nil {
		return ids.GetOrganizationOrUserIdentifiers()
	}
	if ids := ids.GetUserIds(); ids != nil {
		return ids.GetOrganizationOrUserIdentifiers()
	}
	return nil
}
