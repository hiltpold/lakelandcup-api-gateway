package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
)

type CreateLeagueRequestBody struct {
	Admin             string `json:"admin"`
	AdminID           string `json:"adminID"`
	Commissioner      string `json:"commissioner"`
	CommissionerID    string `json:"commissionerID"`
	Name              string `json:"name"`
	FoundationYear    string `json:"foundationYear"`
	MaxFranchises     int32  `json:"maxFranchises,string"`
	MaxProspects      int32  `json:"maxProspects,string"`
	DraftRightsGoalie int32  `json:"draftRightsGoalie,string"`
	DraftRightsSkater int32  `json:"draftRightsSkater,string"`
}

type CreateFranchiseRequestBody struct {
	Name           string `json:"name"`
	OwnerID        string `json:"ownerId"`
	OwnerName      string `json:"ownerName"`
	FoundationYear string `json:"foundationYear"`
	LeagueID       string `json:"leagueId"`
}

type CreateProspectsBulkRequestBody struct {
	ID string `json:"id"`
	FullName string `json:"nhlProspectFullName"`
	FirstName string `json:"nhlProspectFirstName"`
	LastName string `json:"nhlProspectLastName"`
	NhlTeam string `json:"nhlTeamName"`
	Birthdate string `json:"nhlBirthdate"`
	Height string `json:"nhlHeight"`
	Weight int `json:"nhlWeight"`
	NhlDraftYear int `json:"nhlYear"`
	NhlDraftRound string `json:"nhlRound"`
	NhlDraftPickOverall int `json:"nhlPickOverall"`
	NhlDraftPickInRound int `json:"nhlPickInRound"`
	NhlPositionCode string `json:"nhlPositionCode"`
	LeagueID string `json:"leagueId"`
	FranchiseID string `json:"franchiseId"`
	PickID string `json:"pickId"`
}


func CreateLeague(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := CreateLeagueRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateLeague(context.Background(), &pb.LeagueRequest{
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
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func CreateFranchise(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := CreateFranchiseRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateFranchise(context.Background(), &pb.FranchiseRequest{
		Name:           b.Name,
		OwnerID:        b.OwnerID,
		OwnerName:      b.OwnerName,
		FoundationYear: b.FoundationYear,
		LeagueId:       b.LeagueID,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

func CreateProspectsBulk(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := []CreateProspectsBulkRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req := pb.CreateProspectsBulkRequest{}
	for _, p := range(b){
		tmp := pb.CreateProspect{
			FullName: p.FullName,
			FirstName: p.FirstName,
			LastName: p.LastName,
			Birthdate: p.Birthdate,
			Height: p.Height,
			Weight: fmt.Sprintf("%v",p.Weight),
			NhlTeam: p.NhlTeam,
			DraftYear: fmt.Sprintf("%v",p.NhlDraftYear),
			NhlDraftRound: p.NhlDraftRound,
			NhlDraftPickInRound: fmt.Sprintf("%v",p.NhlDraftPickInRound),
			NhlDraftPickOverall: fmt.Sprintf("%v",p.NhlDraftPickOverall),
			PositionCode: p.NhlPositionCode,
		}
		req.Prospects = append(req.Prospects, &tmp)
	}
	
	res, err := c.CreateProspectsBulk(context.Background(), &req)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	

	ctx.JSON(http.StatusCreated, &res)
}
