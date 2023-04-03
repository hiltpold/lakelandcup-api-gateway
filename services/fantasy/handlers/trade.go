package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
	"github.com/sirupsen/logrus"
)

type TradeRequest struct {
	First  *pb.TradePayload `json:"First" binding:"required"`
	Second *pb.TradePayload `json:"Second" binding:"required"`
}

func Trade(ctx *gin.Context, c pb.FantasyServiceClient) {
	b := TradeRequest{}
	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logrus.Info(fmt.Sprintf("%+v", b))

	res, err := c.Trade(context.Background(), &pb.TradeRequest{First: b.First, Second: b.Second})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
