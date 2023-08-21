package product

import (
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/product/pb"
	"log"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	return pb.NewProductServiceClient(cc)
}
