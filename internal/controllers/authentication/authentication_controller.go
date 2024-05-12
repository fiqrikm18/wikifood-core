package authentication

import (
	"WikfoodCore/internal/infrastuctures/persistence"
	"WikfoodCore/internal/utils/jwt"
	"context"
	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	Ctx          context.Context
	DbConnection *persistence.DBConnection
	TokenUtils   *jwt.JWTToken
}

func New(ctx context.Context, tokenUtils *jwt.JWTToken) *AuthenticationController {
	return &AuthenticationController{
		Ctx:        ctx,
		TokenUtils: tokenUtils,
	}
}

// Login
// @BasePath	/api/v1
//
//	@Summary	Login
//	@Schemes
//	@Description	Get Token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	http.SuccessResponse
//	@Router			/auth/login [post]
func (controller *AuthenticationController) Login(g *gin.Context) {
}

// Register
// @BasePath	/api/v1
//
//	@Summary	Register
//	@Schemes
//	@Description	Register New User
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	http.SuccessResponse
//	@Router			/auth/register [post]
func (controller *AuthenticationController) Register(g *gin.Context) {

}

// Logout
// @BasePath	/api/v1
//
//	@Summary	Logout
//	@Schemes
//	@Description	Revoke Token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	http.SuccessResponse
//	@Router			/auth/logout [post]
func (controller *AuthenticationController) Logout(g *gin.Context) {

}

// Refresh Token
// @BasePath	/api/v1
//
//	@Summary	Refresh Token
//	@Schemes
//	@Description	Get New Token
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	http.SuccessResponse
//	@Router			/auth/refresh-token [post]
func (controller *AuthenticationController) RefreshToken(g *gin.Context) {

}
