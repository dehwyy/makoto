package handler

import (
	"context"

	"github.com/dehwyy/Makoto/backend/distributor/graphql/model"
	"github.com/dehwyy/Makoto/backend/distributor/middleware"
)

// shortcut for "auth function(s)"
type _auth_f struct{}

var (
	authf = new(_auth_f)
)

func (h *_auth_f) createAuthResponse(userId, a_token, r_token string) *model.UserAuthResponse {
	return &model.UserAuthResponse{
		UserID: userId,
		Tokens: &model.Tokens{
			AccessToken:  a_token,
			RefreshToken: r_token,
		},
	}
}

func (h *_auth_f) createAuthResponseByCtx(ctx context.Context) *model.UserAuthResponse {
	auth_ctx := middleware.ReadAuthContext(ctx)
	return h.createAuthResponse(auth_ctx.UserId, auth_ctx.AccessToken, auth_ctx.RefreshToken)
}
