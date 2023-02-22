package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
)

func UserInfo(ctx *gin.Context, c pb.AuthServiceClient) {
	val, _ := ctx.Get("userId")
	ID, ok := val.(string)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	val, _ = ctx.Get("role")
	role, ok := val.(string)

	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusCreated, &pb.ValidateResponse{Status: http.StatusCreated, UserId: ID, Role: role})
}
