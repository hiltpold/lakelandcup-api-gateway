package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
	"github.com/sirupsen/logrus"
)

type UpdateLeagueRequestBody struct {
	Admin             string `json:"admin"`
	AdminID           string `json:"adminID"`
	Commissioner      string `json:"commissioner"`
	CommissionerID    string `json:"commissionerID"`
	Name              string `json:"name"`
	FoundationYear    string `json:"foundationYear"`
	MaxFranchises     int32  `json:"maxFranchises"`
	MaxProspects      int32  `json:"maxProspects"`
	DraftRightsGoalie int32  `json:"draftRightsGoalie"`
	DraftRightsSkater int32  `json:"draftRightsSkater"`
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
	logrus.Info(b)

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
	}

	res, err := c.UpdateLeague(context.Background(), &pb.LeagueUpdateRequest{Id: id, League: lR})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
