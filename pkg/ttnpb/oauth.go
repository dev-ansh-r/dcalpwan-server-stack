package ttnpb

// All EntityType methods implement the IDStringer interface.

func (m *OAuthClientAuthorizationIdentifiers) EntityType() string {
	return m.GetUserIds().EntityType()
}

func (m *OAuthClientAuthorization) EntityType() string {
	return m.GetUserIds().EntityType()
}

func (m *ListOAuthClientAuthorizationsRequest) EntityType() string {
	return m.GetUserIds().EntityType()
}

func (m *OAuthAuthorizationCode) EntityType() string {
	return m.GetUserIds().EntityType()
}

func (m *OAuthAccessTokenIdentifiers) EntityType() string {
	return m.GetUserIds().EntityType()
}

func (m *OAuthAccessToken) EntityType() string {
	return m.GetUserIds().EntityType()
}

func (m *ListOAuthAccessTokensRequest) EntityType() string {
	return m.GetUserIds().EntityType()
}

// All IDString methods implement the IDStringer interface.

func (m *OAuthClientAuthorizationIdentifiers) IDString() string {
	return m.GetUserIds().IDString()
}

func (m *OAuthClientAuthorization) IDString() string {
	return m.GetUserIds().IDString()
}

func (m *ListOAuthClientAuthorizationsRequest) IDString() string {
	return m.GetUserIds().IDString()
}

func (m *OAuthAuthorizationCode) IDString() string {
	return m.GetUserIds().IDString()
}

func (m *OAuthAccessTokenIdentifiers) IDString() string {
	return m.GetUserIds().IDString()
}

func (m *OAuthAccessToken) IDString() string {
	return m.GetUserIds().IDString()
}

func (m *ListOAuthAccessTokensRequest) IDString() string {
	return m.GetUserIds().IDString()
}

// All ExtractRequestFields methods are used by github.com/grpc-ecosystem/go-grpc-middleware/tags.

func (m *OAuthClientAuthorizationIdentifiers) ExtractRequestFields(dst map[string]interface{}) {
	m.GetUserIds().ExtractRequestFields(dst)
}

func (m *OAuthClientAuthorization) ExtractRequestFields(dst map[string]interface{}) {
	m.GetUserIds().ExtractRequestFields(dst)
}

func (m *ListOAuthClientAuthorizationsRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetUserIds().ExtractRequestFields(dst)
}

func (m *OAuthAuthorizationCode) ExtractRequestFields(dst map[string]interface{}) {
	m.GetUserIds().ExtractRequestFields(dst)
}

func (m *OAuthAccessTokenIdentifiers) ExtractRequestFields(dst map[string]interface{}) {
	m.GetUserIds().ExtractRequestFields(dst)
}

func (m *OAuthAccessToken) ExtractRequestFields(dst map[string]interface{}) {
	m.GetUserIds().ExtractRequestFields(dst)
}

func (m *ListOAuthAccessTokensRequest) ExtractRequestFields(dst map[string]interface{}) {
	m.GetUserIds().ExtractRequestFields(dst)
}
