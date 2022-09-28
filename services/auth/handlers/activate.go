package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
)

func Activate(ctx *gin.Context, c pb.AuthServiceClient) {

	token := ctx.Param("token")

	res, err := c.Activate(context.Background(), &pb.ActivateRequest{
		Token: token,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
