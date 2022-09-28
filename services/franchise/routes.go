package franchise

import (
	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/config"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth"
	handler "github.com/hiltpold/lakelandcup-api-gateway/services/franchise/handlers"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	versionGroup := r.Group("/v1")
	routes := versionGroup.Group("/order")

	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	handler.CreateFranchise(ctx, svc.Client)
}
