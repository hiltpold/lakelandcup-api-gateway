package fantasy

import (
	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/config"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth"
	handler "github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/handlers"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	//a := auth.InitAuthMiddleware(authSvc)
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	fantasyGroup := r.Group("/v1/fantasy/")

	//fantasyGroup.Use(a.AuthRequired)
	// league routes
	fantasyGroup.POST("/league", svc.CreateLeague)
	fantasyGroup.GET("/leagues", svc.GetLeagues)
	fantasyGroup.POST("/league/:id", svc.UpdateLeague)
	fantasyGroup.GET("/league/:id/franchises", svc.GetLeagueFranchises)
	// franchise routes
	fantasyGroup.POST("/franchise", svc.CreateFranchise)
	fantasyGroup.GET("/franchise/:id", svc.GetFranchise)
	// prospect routes
	fantasyGroup.POST("/prospects", svc.CreateProspectsBulk)
	fantasyGroup.GET("/prospects/:text", svc.TextSearchProspects)
}

func (svc *ServiceClient) CreateLeague(ctx *gin.Context) {
	handler.CreateLeague(ctx, svc.Client)
}

func (svc *ServiceClient) UpdateLeague(ctx *gin.Context) {
	handler.UpdateLeague(ctx, svc.Client)
}

func (svc *ServiceClient) GetLeagues(ctx *gin.Context) {
	handler.GetLeagues(ctx, svc.Client)
}

func (svc *ServiceClient) GetLeagueFranchises(ctx *gin.Context) {
	handler.GetLeagueFranchises(ctx, svc.Client)
}

func (svc *ServiceClient) CreateFranchise(ctx *gin.Context) {
	handler.CreateFranchise(ctx, svc.Client)
}

func (svc *ServiceClient) GetFranchise(ctx *gin.Context) {
	handler.GetFranchise(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProspectsBulk(ctx *gin.Context) {
	handler.CreateProspectsBulk(ctx, svc.Client)
}

func (svc *ServiceClient) TextSearchProspects(ctx *gin.Context) {
	handler.TextSearchProspects(ctx, svc.Client)
}
