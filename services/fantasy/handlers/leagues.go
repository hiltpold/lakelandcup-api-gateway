package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
)

func GetLeagueFranchises(ctx *gin.Context, c pb.FantasyServiceClient) {
	id := ctx.Param("id")

	res, err := c.GetLeagueFranchises(context.Background(), &pb.GetLeagueRequest{LeagueId: id})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func GetLeagues(ctx *gin.Context, c pb.FantasyServiceClient) {

	res, err := c.GetLeagues(context.Background(), &pb.GetLeaguesRequest{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
