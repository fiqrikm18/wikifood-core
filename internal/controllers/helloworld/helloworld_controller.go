package helloworld

import (
	"WikfoodCore/internal/dto/helloworld"
	httpUtils "WikfoodCore/internal/utils/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloWorldController struct{}

func New() *HelloWorldController {
	return &HelloWorldController{}
}

// Ping
// @BasePath	/api/v1
//
//	@Summary	ping
//	@Schemes
//	@Description	do ping
//	@Tags			Ping
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	http.SuccessResponse{data=helloworld.HelloworldResponse}
//	@Router			/helloworld [get]
func (controller *HelloWorldController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, httpUtils.NewSuccessResponse(http.StatusOK, "Success", helloworld.HelloworldResponse{
		Message: "helloworld",
	}))
}
