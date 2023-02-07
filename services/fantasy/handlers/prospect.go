package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
)

type DraftPick struct {
	DraftYear        string `json:"draftYear"`
	DraftRound       string `json:"draftRound"`
	DraftPickInRound string `json:"draftPickInRound"`
	DraftPickOverall string `json:"draftPickOverall"`
}

type DraftProspectRequest struct {
	DraftPick   DraftPick `json:"draftPick" binding:"required"`
	LeagueID    string    `json:"leagueID"`
	FranchiseID string    `json:"franchiseID"`
	ProspectID  string    `json:"prospectID"`
}

func DraftProspect(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := DraftProspectRequest{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	dp := pb.DraftPick{DraftYear: b.DraftPick.DraftYear, DraftRound: b.DraftPick.DraftRound, DraftPickInRound: b.DraftPick.DraftPickInRound, DraftPickOverall: b.DraftPick.DraftPickOverall}
	res, err := c.DraftProspect(context.Background(), &pb.DraftProspectRequest{DraftPick: &dp, LeagueID: b.LeagueID, FranchiseID: b.FranchiseID, ProspectID: b.ProspectID})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
