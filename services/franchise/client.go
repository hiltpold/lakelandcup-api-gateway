package franchise

import (
	"fmt"

	"github.com/hiltpold/lakelandcup-api-gateway/config"
	"github.com/hiltpold/lakelandcup-api-gateway/services/franchise/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.FranchiseServiceClient
}

func InitServiceClient(c *config.Config) pb.FranchiseServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.FranchiseSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewFranchiseServiceClient(cc)
}
