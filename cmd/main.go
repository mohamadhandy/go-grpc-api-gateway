package main

import (
	"go-grpc-api-gateway/pkg/auth"
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/order"
	"go-grpc-api-gateway/pkg/product"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed at config: %v", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
