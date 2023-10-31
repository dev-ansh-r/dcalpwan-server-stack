package oauth

import (
	"github.com/openshift/osin"
	"go.thethings.network/lorawan-stack/v3/pkg/auth"
)

func (s *server) GenerateAuthorizeToken(_ *osin.AuthorizeData) (string, error) {
	return auth.AuthorizationCode.Generate(s.c.Context(), "")
}

func (s *server) GenerateAccessToken(_ *osin.AccessData, generateRefresh bool) (accessToken string, refreshToken string, err error) {
	ctx := s.c.Context()
	var id string
	if generateRefresh {
		id, err = auth.GenerateID(ctx)
		if err != nil {
			return "", "", err
		}
	}
	accessToken, err = auth.AccessToken.Generate(ctx, id)
	if err != nil {
		return "", "", err
	}
	if generateRefresh {
		refreshToken, err = auth.RefreshToken.Generate(ctx, id)
		if err != nil {
			return "", "", err
		}
	}
	return accessToken, refreshToken, nil
}
