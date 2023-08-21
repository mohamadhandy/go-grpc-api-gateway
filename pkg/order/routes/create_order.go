package routes

import (
	"go-grpc-api-gateway/pkg/order/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	bodyOrder := OrderRequestBody{}
	if err := ctx.BindJSON(&bodyOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")
	res, err := c.CreateOrder(ctx, &pb.CreateOrderRequest{
		ProductId: bodyOrder.ProductId,
		Quantity:  bodyOrder.Quantity,
		UserId:    userId.(int64),
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
