package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
)

type ResetPasswordRequestBody struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func ResetPassword(ctx *gin.Context, c pb.AuthServiceClient) {
	token := ctx.Param("token")

	b := ResetPasswordRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ResetPassword(context.Background(), &pb.ResetPasswordRequest{
		Token:           token,
		Password:        b.Password,
		ConfirmPassword: b.ConfirmPassword,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
