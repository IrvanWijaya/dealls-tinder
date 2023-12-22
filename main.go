package main

import (
	"log"

	"github.com/IrvanWijaya/dealls-tinder/controllers"
	"github.com/IrvanWijaya/dealls-tinder/initializers"
	"github.com/IrvanWijaya/dealls-tinder/routes"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController()
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	router := server.Group("/api")
	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
