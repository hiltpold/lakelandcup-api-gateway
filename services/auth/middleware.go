package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
	"github.com/sirupsen/logrus"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	access_token, _ := ctx.Cookie("lakelandcup_access_token")

	res, err := c.svc.Client.Validate(ctx, &pb.ValidateRequest{
		Token:     access_token,
		TokenType: "ACCESS_TOKEN",
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "error": res.Error})
		return
	}
	logrus.Info(res)
	ctx.Set("userId", res.UserId)
	ctx.Set("role", res.Role)

	ctx.Next()
}
