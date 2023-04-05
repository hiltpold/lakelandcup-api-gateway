package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/conf"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth/pb"
)

func SignOut(ctx *gin.Context, c pb.AuthServiceClient, config *conf.Configuration) {

	ctx.SetCookie("lakelandcup_access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("lakelandcup_refresh_token", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
