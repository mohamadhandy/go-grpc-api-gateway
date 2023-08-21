package auth

import (
	"go-grpc-api-gateway/pkg/auth/pb"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")
	if len(token) > 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	res, err := c.svc.Client.Validate(ctx, &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil && res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set("userId", res.UserId)
	ctx.Next()
}
