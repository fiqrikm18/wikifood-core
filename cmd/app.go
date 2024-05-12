package cmd

import (
	"WikfoodCore/internal/routes"
	"github.com/gin-gonic/gin"
	"os"

	"WikfoodCore/docs"
)

//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	fiqrikm18@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func RunServer() {
	setupSwaggerInfo()

	r := gin.Default()
	r.Use(gin.Recovery())

	webRoute := routes.NewWebRoute()
	apiRoute := routes.NewAPIRoute()

	webRoute.Handle(r)
	apiRoute.Handle(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func setupSwaggerInfo() {
	docs.SwaggerInfo.Title = "Wikifood Core"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is a sample recipe sharing API."
	docs.SwaggerInfo.Host = os.Getenv("APP_BASEURL")

	docs.SwaggerInfo.BasePath = "/api/v1"
}
