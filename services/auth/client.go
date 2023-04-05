package auth

import (
	"fmt"

	"github.com/hiltpold/lakelandcup-api-gateway/conf"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *conf.Configuration) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
