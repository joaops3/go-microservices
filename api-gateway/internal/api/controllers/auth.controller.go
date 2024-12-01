package controllers

import (
	"go-microservices-grpc/api-gateway/internal/api/dtos"
	"go-microservices-grpc/api-gateway/internal/pb"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService pb.AuthServiceClient
}

func InitAuthController(authService pb.AuthServiceClient) *AuthController {
	controller := &AuthController{AuthService: authService}
	return controller
}

func (c *AuthController) SignIn(ctx *gin.Context){

	dto := dtos.SignInDto{}

    err := ctx.BindJSON(&dto)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    err = dto.Validate()
    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    data, err := c.AuthService.SignIn(ctx, dto.ToProtoBuff())

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    sendSuccess(ctx, "success", data)
    return
}


func (c *AuthController) SignUp(ctx *gin.Context){

	dto := dtos.CreateUserDto{}

    err := ctx.BindJSON(&dto)

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    err = dto.Validate()
    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

    data, err := c.AuthService.SignUp(ctx, dto.ToProtoBuff())

    if err != nil {
       sendError(ctx, 400, err.Error())
        return 
    }

   sendSuccess(ctx, "success", data)
    return
}