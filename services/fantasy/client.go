package fantasy

import (
	"fmt"

	"github.com/hiltpold/lakelandcup-api-gateway/config"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.FantasyServiceClient
}

func InitServiceClient(c *config.Config) pb.FantasyServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.FranchiseSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewFantasyServiceClient(cc)
}
