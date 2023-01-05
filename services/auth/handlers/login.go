package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/config"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
	"github.com/sirupsen/logrus"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient, config *config.Config) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	//ctx.SetSameSite(http.SameSiteLaxMode)
	//TODO: better configurability, especially for productional system
	ctx.SetCookie("lakelandcup_refresh_token", res.RefreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("lakelandcup_access_token", res.Token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusCreated, &res)
	logrus.Info(fmt.Sprintf("%v", ctx.Request.Cookies()))
}
