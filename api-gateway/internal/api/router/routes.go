package router

import (
	"go-microservices-grpc/api-gateway/internal/api/client"
	"go-microservices-grpc/api-gateway/internal/api/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	

	InitializeAuthRoutes(r)
	
}

func InitializeAuthRoutes(r *gin.Engine){
	config := &client.Config{
		Port:      os.Getenv("AUTH_PORT"),
		AuthSuvUrl: os.Getenv("AUTH_SVC_URL"),
	}

	c := client.InitServiceClient(config)
	controller := controllers.InitAuthController(c)

	routerGroup := r.Group("/auth")
	routerGroup.POST("/signin", controller.SignIn)
	routerGroup.POST("/signup",controller.SignUp)
}


