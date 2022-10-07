package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
)

type ActivationRequestBody struct {
	Email string `json:"email"`
}

func ResendActivationToken(ctx *gin.Context, c pb.AuthServiceClient) {
	b := ActivationRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.ResendActivationToken(context.Background(), &pb.ResendActivationTokenRequest{
		Email: b.Email,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
