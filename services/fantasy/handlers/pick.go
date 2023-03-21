package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
)

type CreateOrUpdatePick struct {
	Franchise       string `json:"Franchise"`
	FranchiseID     string `json:"FranchiseID"`
	OwnerName       string `json:"Owner"`
	Year            string `json:"Year"`
	LotteryPosition int    `json:"LotteryPosition"`
}

type CreateOrUpdatePicksRequestBody struct {
	LeagueID string                   `json:"LeagueID"`
	Picks    []*pb.CreateOrUpdatePick `json:"Picks" binding:"dive"`
}

func CreateOrUpdatePicks(ctx *gin.Context, c pb.FantasyServiceClient) {
	id := ctx.Param("id")
	b := CreateOrUpdatePicksRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req := &pb.CreateOrUpdatePicksRequest{
		LeagueID: id,
		Picks:    b.Picks,
	}

	res, err := c.CreateOrUpdatePicks(context.Background(), req)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
