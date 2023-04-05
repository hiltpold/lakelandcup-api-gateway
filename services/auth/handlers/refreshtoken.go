package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/conf"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
)

type RefreshTokenRequestBody struct {
	string `json:"refresh_token"`
}

func RefreshToken(ctx *gin.Context, c pb.AuthServiceClient, config *conf.Configuration) {
	refreshToken, getRefreshTokenErr := ctx.Cookie("lakelandcup_refresh_token")

	if getRefreshTokenErr != nil {
		ctx.AbortWithError(http.StatusBadGateway, getRefreshTokenErr)
		return
	}

	b := RefreshTokenRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.RefreshToken(context.Background(), &pb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
