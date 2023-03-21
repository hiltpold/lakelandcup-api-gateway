package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
	"github.com/sirupsen/logrus"
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
	FullName    string `json:"FullName"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	NhlTeam     string `json:"NhlTeam"`
	Birthdate   string `json:"Birthdate"`
	LeagueID    string `json:"LeagueID"`
	FranchiseID string `json:"FranchiseID"`
}
type TradeRequest struct {
	FromFranchiseID string     `json:"fromLeagueID"`
	ToFranchiseID   string     `json:"toLeagueID"`
	Picks           []Pick     `json:"picks" binding:"required"`
	Prospects       []Prospect `json:"prospects" binding:"required"`
}

func Trade(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := TradeRequest{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logrus.Info(fmt.Sprintf("%v", b))

	picks := []*pb.Pick{}
	prospects := []*pb.Prospect{}

	for _, p := range b.Picks {
		picks = append(picks, &pb.Pick{
			DraftYear:        p.DraftYear,
			DraftRound:       p.DraftRound,
			DraftPickInRound: p.DraftPickInRound,
			DraftPickOverall: p.DraftPickOverall,
			ProspectID:       p.ProspectID,
			OwnerID:          p.OwnerID,
			OwnerName:        p.OwnerName,
			LastOwnerID:      p.LastOwnerID,
			LastOwnerName:    p.LastOwnerName,
			OriginID:         p.OriginID,
		})
	}

	for _, p := range b.Prospects {
		prospects = append(prospects, &pb.Prospect{
			FullName:    p.FullName,
			FirstName:   p.FirstName,
			LastName:    p.LastName,
			NhlTeam:     p.NhlTeam,
			Birthdate:   p.Birthdate,
			LeagueID:    p.LeagueID,
			FranchiseID: p.FranchiseID})
	}

	res, err := c.Trade(context.Background(), &pb.TradeRequest{FromFranchiseID: b.FromFranchiseID, ToFranchiseID: b.ToFranchiseID, Picks: picks, Prospects: prospects})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
