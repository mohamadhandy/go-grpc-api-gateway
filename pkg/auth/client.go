package auth

import (
	"go-grpc-api-gateway/pkg/auth/pb"
	"go-grpc-api-gateway/pkg/config"
	"log"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// with insecure, because no ssl running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	return pb.NewAuthServiceClient(cc)
}
