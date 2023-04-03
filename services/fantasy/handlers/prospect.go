package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
)

type DraftProspectRequest struct {
	LeagueID    string `json:"LeagueID"`
	FranchiseID string `json:"FranchiseID"`
	ProspectID  string `json:"ProspectID"`
	PickID      string `json:"PickID"`
}

func DraftProspect(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := DraftProspectRequest{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.DraftProspect(context.Background(), &pb.DraftRequest{PickID: b.PickID, LeagueID: b.LeagueID, FranchiseID: b.FranchiseID, ProspectID: b.ProspectID})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func UndraftProspect(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := DraftProspectRequest{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.UndraftProspect(context.Background(), &pb.DraftRequest{PickID: b.PickID, LeagueID: b.LeagueID, FranchiseID: b.FranchiseID, ProspectID: b.ProspectID})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func GetProspectsByFranchise(ctx *gin.Context, c pb.FantasyServiceClient) {
	id := ctx.Param("id")
	year := ctx.Param("franchise")

	req := &pb.GetFranchiseRequest{
		LeagueID:    id,
		FranchiseID: year,
	}

	res, err := c.GetProspectsByFranchise(context.Background(), req)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
