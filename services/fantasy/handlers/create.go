package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
)

type CreateLeagueRequestBody struct {
	UserId         string `json:"userId"`
	LeagueName     string `json:"leagueName"`
	FoundationYear string `json:"foundationYear"`
}

func CreateLeague(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := CreateLeagueRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateLeague(context.Background(), &pb.LeagueRequest{
		UserId:         b.UserId,
		LeagueName:     b.LeagueName,
		FoundationYear: b.FoundationYear,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
