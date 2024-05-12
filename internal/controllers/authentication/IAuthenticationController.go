package authentication

import "github.com/gin-gonic/gin"

type IAuthenticationController interface {
	Login(g *gin.Context)
	Register(g *gin.Context)
	Logout(g *gin.Context)
	RefreshToken(g *gin.Context)
}
