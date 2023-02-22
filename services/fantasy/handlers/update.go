package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
)

type UpdateLeagueRequestBody struct {
	Admin             string `json:"Admin"`
	AdminID           string `json:"AdminID"`
	Commissioner      string `json:"Commissioner"`
	CommissionerID    string `json:"CommissionerID"`
	Name              string `json:"Name"`
	FoundationYear    string `json:"FoundationYear"`
	MaxFranchises     int32  `json:"MaxFranchises"`
	MaxProspects      int32  `json:"MaxProspects"`
	DraftRightsGoalie int32  `json:"DraftRightsGoalie"`
	DraftRightsSkater int32  `json:"DraftRightsSkater"`
	DraftRounds       int32  `json:"DraftRounds"`
}

type UpdateFranchiseRequestBody struct {
	Name           string `json:"name"`
	OwnerID        string `json:"ownerId"`
	OwnerName      string `json:"ownerName"`
	FoundationYear string `json:"foundationYear"`
	LeagueID       string `json:"leagueId"`
}

func UpdateLeague(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := UpdateLeagueRequestBody{}
	id := ctx.Param("id")

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	lR := &pb.LeagueRequest{
		Name:              b.Name,
		Admin:             b.Admin,
		AdminID:           b.AdminID,
		Commissioner:      b.Commissioner,
		CommissionerID:    b.CommissionerID,
		FoundationYear:    b.FoundationYear,
		MaxFranchises:     b.MaxFranchises,
		MaxProspects:      b.MaxProspects,
		DraftRightsGoalie: b.DraftRightsGoalie,
		DraftRightsSkater: b.DraftRightsSkater,
		DraftRounds:       b.DraftRounds,
	}

	res, err := c.UpdateLeague(context.Background(), &pb.LeagueUpdateRequest{Id: id, League: lR})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
