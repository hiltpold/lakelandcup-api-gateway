package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/conf"
	handler "github.com/hiltpold/lakelandcup-api-gateway/services/auth/handlers"
)

func RegisterProtectedRoutes(r *gin.Engine, c *conf.Configuration, authSvc *ServiceClient) {
	a := InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	protected := r.Group("/v1/auth/user/")
	protected.Use(a.AuthRequired)
	protected.GET("/info", svc.UserInfo)
	protected.GET("/all", svc.GetUsers)
}

func RegisterRoutes(r *gin.Engine, c *conf.Configuration) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	version := r.Group("/v1")
	authRoutes := version.Group("/auth")

	authRoutes.POST("/signup", svc.Register)
	authRoutes.POST("/signin", svc.Login(c))
	authRoutes.GET("/signout", svc.SignOut(c))
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

func (svc *ServiceClient) Login(c *conf.Configuration) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		handler.Login(ctx, svc.Client, c)
	}
	return fn
}

func (svc *ServiceClient) SignOut(c *conf.Configuration) gin.HandlerFunc {
	fn := func(ctx *gin.Context) {
		handler.SignOut(ctx, svc.Client, c)
	}
	return fn
}

func (svc *ServiceClient) RefreshToken(c *conf.Configuration) gin.HandlerFunc {
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

func (svc *ServiceClient) GetUsers(ctx *gin.Context) {
	handler.GetUsers(ctx, svc.Client)
}

func (svc *ServiceClient) UserInfo(ctx *gin.Context) {
	handler.UserInfo(ctx, svc.Client)
}
