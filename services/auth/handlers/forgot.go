package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
)

type ForgotPasswordRequestBody struct {
	Email string `json:"email"`
}

func ForgotPassword(ctx *gin.Context, c pb.AuthServiceClient) {
	b := ForgotPasswordRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ForgotPassword(context.Background(), &pb.ForgotPasswordRequest{
		Email: b.Email,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
