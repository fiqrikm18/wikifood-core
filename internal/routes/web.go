package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WebRoute struct {
}

func NewWebRoute() *WebRoute {
	return &WebRoute{}
}

func (route *WebRoute) Handle(eng *gin.Engine) {
	// swagger route
	eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
