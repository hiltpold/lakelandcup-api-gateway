package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/config"
	handler "github.com/hiltpold/lakelandcup-api-gateway/services/auth/handlers"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	version := r.Group(("/v1"))
	routes := version.Group("/auth")

	routes.POST("/register", svc.Register)
	routes.POST("/login", svc.Login)
	routes.POST("/activation/:token", svc.Activation)
	routes.POST("/activation/token/resend", svc.ResendActivationToken)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	handler.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	handler.Login(ctx, svc.Client)
}

func (svc *ServiceClient) Activation(ctx *gin.Context) {
	handler.Activate(ctx, svc.Client)
}

func (svc *ServiceClient) ResendActivationToken(ctx *gin.Context) {
	handler.Activate(ctx, svc.Client)
}
