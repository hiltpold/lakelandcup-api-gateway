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
	authRoutes.POST("/user/signin", svc.Login(c))
	authRoutes.GET("/user", svc.UserInfo)
	authRoutes.POST("/refreshtoken", svc.RefreshToken(c))
	authRoutes.GET("/account/activation/:token", svc.Activation)
	authRoutes.POST("/account/activation/token/resend", svc.ResendActivationToken)
	authRoutes.PUT("/account/password", svc.ForgotPassword)
	authRoutes.PUT("/account/password/reset", svc.ResetPassword)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	handler.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(c *config.Config) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		handler.Login(ctx, svc.Client, c)
	}
	return fn
}

func (svc *ServiceClient) UserInfo(ctx *gin.Context) {
	handler.UserInfo(ctx, svc.Client)
}

func (svc *ServiceClient) RefreshToken(c *config.Config) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		handler.RefreshToken(ctx, svc.Client, c)
	}
	return fn
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
