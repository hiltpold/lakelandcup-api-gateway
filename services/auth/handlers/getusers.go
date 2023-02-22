package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
)

func GetUsers(ctx *gin.Context, c pb.AuthServiceClient) {
	val, _ := ctx.Get("userId")
	userId, ok := val.(string)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	res, err := c.GetUsers(context.Background(), &pb.GetUsersRequest{UserID: string(userId)})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": err})
		return
	}
	ctx.JSON(http.StatusCreated, &res)
}
