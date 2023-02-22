package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
)

type Pick struct {
	DraftYear        string `json:"DraftYear"`
	DraftRound       string `json:"DraftRound"`
	DraftPickInRound string `json:"DraftPickInRound"`
	DraftPickOverall string `json:"DraftPickOverall"`
	ProspectID       string `json:"ProspectID"`
	OwnerID          string `json:"OwnerID"`
	OwnerName        string `json:"OwnerName"`
	LastOwnerID      string `json:"LastOwnerID"`
	LastOwnerName    string `json:"LastOwnerName"`
	OriginID         string `json:"OriginID"`
	OriginName       string `json:"OriginName"`
}
type Prospect struct {
	FullName      string `json:"FullName"`
	NhlTeam       string `json:"NhlTeam"`
	Birthdate     string `json:"Birthdate"`
	ProspectID    string `json:"ProspectID"`
	LeagueID      string `json:"LeagueID`
	OwnerID       string `json:"OwnerID"`
	OwnerName     string `json:"OwnerName"`
	LastOwnerID   string `json:"LastOwnerID"`
	LastOwnerName string `json:"LastOwnerName"`
	OriginID      string `json:"OriginID"`
	Origin        string `json:"Origin"`
}
type TradeRequest struct {
	FromFranchiseID string     `json:"fromLeagueID"`
	ToFranchiseID   string     `json:"toLeagueID"`
	Pick            []Pick     `json:"picks" binding:"required"`
	Prospects       []Prospect `json:"prospects" binding:"required"`
}

func Trade(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := TradeRequest{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Trade(context.Background(), &pb.TradeRequest{FranchiseID: id})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
