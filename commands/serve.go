package commands

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/conf"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy"
	"github.com/spf13/cobra"
)

var serveCmd = cobra.Command{
	Use:  "serve",
	Long: "Start Fantasy Service",
	Run: func(cmd *cobra.Command, args []string) {
		runWithConfig(cmd, serve)
	},
}

func serve(c *conf.Configuration) {

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowCredentials = true
	//config.AllowMethods = []string{"GET", "POST"}

	r.Use(cors.New(config))

	authSvc := *auth.RegisterRoutes(r, c)
	auth.RegisterProtectedRoutes(r, c, &authSvc)
	fantasy.RegisterRoutes(r, c, &authSvc)

	r.Run(":" + c.Port)
}
