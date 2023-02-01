package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/pb"
	"github.com/sirupsen/logrus"
)

func TextSearchProspects(ctx *gin.Context, c pb.FantasyServiceClient) {
	ft := ctx.Param("text")
	logrus.Info(fmt.Sprintf("Full Text: %v", ft))
	res, err := c.TextSearchProspects(context.Background(), &pb.TextSearchRequest{
		Text: ft,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
