package routes

import (
	"WikfoodCore/internal/controllers/authentication"
	"WikfoodCore/internal/controllers/helloworld"
	"WikfoodCore/internal/infrastuctures/persistence"
	"WikfoodCore/internal/middlewares"
	TokenRepository "WikfoodCore/internal/repositories/token"
	UserRepository "WikfoodCore/internal/repositories/user"
	TokenService "WikfoodCore/internal/services/token"
	UserService "WikfoodCore/internal/services/user"
	jwtUtils "WikfoodCore/internal/utils/jwt"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	helloWorldController     helloworld.IHelloWorldController
	authenticationController authentication.IAuthenticationController
)

var (
	HELLOWORLD_PATH     = "/helloworld"
	AUTHENTICATION_PATH = "/auth"
)

var (
	userRepository  UserRepository.IUserRepository
	tokenRepository TokenRepository.ITokenRepository
)

var (
	userService  UserService.IUserService
	tokenService TokenService.ITokenService
)

type APIRoute struct {
	dbPersistence *persistence.DBConnection
}

func NewAPIRoute() *APIRoute {
	dbConnection := persistence.NewDBConnection()
	return &APIRoute{
		dbPersistence: dbConnection,
	}
}

func (route *APIRoute) Handle(eng *gin.Engine) {
	tokenUtils := jwtUtils.NewJWTToken()
	ctx, cancelFn := context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
	defer cancelFn()

	// setup repository
	userRepository = UserRepository.New(ctx, route.dbPersistence)
	tokenRepository = TokenRepository.New(ctx, route.dbPersistence)

	// setup service
	userService = UserService.New(userRepository)
	tokenService = TokenService.New(tokenRepository)

	// setup controller
	helloWorldController = helloworld.New()
	authenticationController = authentication.New(ctx, tokenUtils)

	api := eng.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET(HELLOWORLD_PATH, helloWorldController.Index)

			v1.POST(AUTHENTICATION_PATH+"/login", authenticationController.Login)
			v1.POST(AUTHENTICATION_PATH+"/register", authenticationController.Register)

			v1Auth := v1.Group("")
			{
				v1Auth.Use(middlewares.AuthenticationMiddleware())

				v1Auth.POST(AUTHENTICATION_PATH+"/logout", authenticationController.Logout)
				v1Auth.POST(AUTHENTICATION_PATH+"/refresh-token", authenticationController.RefreshToken)
			}
		}
	}
}
