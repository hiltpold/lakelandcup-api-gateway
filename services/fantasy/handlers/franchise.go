package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
)

func GetFranchise(ctx *gin.Context, c pb.FantasyServiceClient) {
	id := ctx.Param("id")

	res, err := c.GetFranchise(context.Background(), &pb.GetFranchiseRequest{FranchiseId: id})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
