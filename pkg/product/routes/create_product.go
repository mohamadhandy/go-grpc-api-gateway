package routes

import (
	"fmt"
	"go-grpc-api-gateway/pkg/product/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	body := ProductRequestBody{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println("body stock", body.Stock)
	res, err := c.CreateProduct(ctx, &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	ctx.JSON(int(res.Status), &res)
}
