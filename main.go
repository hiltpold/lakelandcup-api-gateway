package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/config"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth"
	"github.com/hiltpold/lakelandcup-api-gateway/services/fantasy"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowCredentials = true
	//config.AllowMethods = []string{"GET", "POST"}

	r.Use(cors.New(config))

	authSvc := *auth.RegisterRoutes(r, &c)
	auth.RegisterProtectedRoutes(r, &c, &authSvc)
	fantasy.RegisterRoutes(r, &c, &authSvc)

	r.Run(":" + c.Port)
}
