package fantasy

import (
	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/conf"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth"
	handler "github.com/hiltpold/lakelandcup-api-gateway/services/fantasy/handlers"
)

func RegisterRoutes(r *gin.Engine, c *conf.Configuration, authSvc *auth.ServiceClient) {
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
	fantasyGroup.POST("/prospect", svc.CreateProspectsBulk)
	fantasyGroup.GET("/prospect/:text", svc.TextSearchProspects)
	// trade between franchises
	fantasyGroup.POST("/league/:id/franchise/trade", svc.Trade)
	// update or create picks according to draft lottery
	fantasyGroup.POST("/league/:id/picks", svc.CreateOrUpdatePicks)
	// get all picks for a given year
	fantasyGroup.GET("/league/:id/picks/:year", svc.GetPicksByYear)
	// get all picks for a given franchise
	fantasyGroup.GET("/league/:id/franchise/:franchise/picks", svc.GetPicksByFranchise)
	// get all prospects for a given franchise
	fantasyGroup.GET("/league/:id/franchise/:franchise/prospects", svc.GetProspectsByFranchise)

	fantasyGroup.POST("/league/:id/draft", svc.DraftProspect)
	fantasyGroup.POST("/league/:id/undraft", svc.UndraftProspect)
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

func (svc *ServiceClient) Trade(ctx *gin.Context) {
	handler.Trade(ctx, svc.Client)
}

func (svc *ServiceClient) CreateOrUpdatePicks(ctx *gin.Context) {
	handler.CreateOrUpdatePicks(ctx, svc.Client)
}

func (svc *ServiceClient) GetPicksByYear(ctx *gin.Context) {
	handler.GetPicksByYear(ctx, svc.Client)
}

func (svc *ServiceClient) GetPicksByFranchise(ctx *gin.Context) {
	handler.GetPicksByFranchise(ctx, svc.Client)
}

func (svc *ServiceClient) GetProspectsByFranchise(ctx *gin.Context) {
	handler.GetProspectsByFranchise(ctx, svc.Client)
}

func (svc *ServiceClient) DraftProspect(ctx *gin.Context) {
	handler.DraftProspect(ctx, svc.Client)
}

func (svc *ServiceClient) UndraftProspect(ctx *gin.Context) {
	handler.UndraftProspect(ctx, svc.Client)
}
