package order

import (
	"go-grpc-api-gateway/pkg/config"
	"go-grpc-api-gateway/pkg/order/pb"
	"log"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	// with insecure, because no ssl running
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	return pb.NewOrderServiceClient(cc)
}
