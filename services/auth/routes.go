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
	//userRoutes := authRoutes.Group("/user")

	authRoutes.POST("/user/signup", svc.Register)
	authRoutes.POST("/user/signin", svc.Login)
	authRoutes.GET("/account/activation/:token", svc.Activation)
	authRoutes.POST("/account/activation/token/resend", svc.ResendActivationToken)
	authRoutes.PUT("/account/password", svc.ForgotPassword)
	authRoutes.PUT("/account/password/reset", svc.ResetPassword)

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
