package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/franchise/pb"
)

type CreateFranchiseRequestBody struct {
	ProductId int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

func CreateFranchise(ctx *gin.Context, c pb.FranchiseServiceClient) {
	b := CreateFranchiseRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	res, err := c.CreateFranchise(context.Background(), &pb.CreateFranchiseRequest{
		ProductId: b.ProductId,
		Quantity:  b.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
