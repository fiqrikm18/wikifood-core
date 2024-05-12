package middlewares

import (
	jwtUtils "WikfoodCore/internal/utils/jwt"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenUtils := jwtUtils.NewJWTToken()
		_, _ = tokenUtils.ValidateToken("", jwtUtils.ACCESS_TOKEN)
		c.Next()
	}
}
