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
	authRoutes := version.Group("/auth")
	userRoutes := authRoutes.Group("/user")

	userRoutes.POST("/register", svc.Register)
	userRoutes.POST("/login", svc.Login)
	userRoutes.POST("/account/activation/:token", svc.Activation)
	userRoutes.POST("/account/activation/token/resend", svc.ResendActivationToken)
	userRoutes.PUT("/account/password", svc.ForgotPassword)
	userRoutes.PUT("/account/password/reset", svc.ResetPassword)

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
	handler.ResendActivationToken(ctx, svc.Client)
}

func (svc *ServiceClient) ForgotPassword(ctx *gin.Context) {
	handler.ForgotPassword(ctx, svc.Client)
}

func (svc *ServiceClient) ResetPassword(ctx *gin.Context) {
	handler.ResetPassword(ctx, svc.Client)
}
