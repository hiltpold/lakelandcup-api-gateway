package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
	"github.com/sirupsen/logrus"
)

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
	logrus.Info(fmt.Sprintf("%v", b))

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

func GetPicksByYear(ctx *gin.Context, c pb.FantasyServiceClient) {
	id := ctx.Param("id")
	year := ctx.Param("year")

	req := &pb.GetPicksRequest{
		LeagueID: id,
		Year:     year,
	}

	res, err := c.GetPicksByYear(context.Background(), req)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func GetPicksByFranchise(ctx *gin.Context, c pb.FantasyServiceClient) {
	id := ctx.Param("id")
	year := ctx.Param("franchise")

	req := &pb.GetPicksRequest{
		LeagueID:    id,
		FranchiseID: year,
	}

	res, err := c.GetPicksByFranchise(context.Background(), req)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
