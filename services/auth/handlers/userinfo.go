package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
)

func UserInfo(ctx *gin.Context, c pb.AuthServiceClient) {
	token, _ := ctx.Cookie("lakelandcup_access_token")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized})
		//ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.Validate(context.Background(), &pb.ValidateRequest{
		Token:     token,
		TokenType: "ACCESS_TOKEN",
	})

	if err != nil || res.Status != http.StatusOK {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized})
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
