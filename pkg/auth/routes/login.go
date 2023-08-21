package routes

import (
	"go-grpc-api-gateway/pkg/auth/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	bodyLogin := LoginRequestBody{}

	if err := ctx.BindJSON(&bodyLogin); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(ctx, &pb.LoginRequest{
		Email:    bodyLogin.Email,
		Password: bodyLogin.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
