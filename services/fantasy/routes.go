package franchise

import (
	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/config"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth"
	handler "github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/handlers"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	fantasyGroup := r.Group("/v1/fantasy")
	leagueRoutes := fantasyGroup.Group("/league")

	leagueRoutes.Use(a.AuthRequired)
	leagueRoutes.POST("/", svc.CreateLeague)
}

func (svc *ServiceClient) CreateLeague(ctx *gin.Context) {
	handler.CreateLeague(ctx, svc.Client)
}
