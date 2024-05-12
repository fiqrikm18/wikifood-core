package helloworld

import "github.com/gin-gonic/gin"

type IHelloWorldController interface {
	Index(c *gin.Context)
}
