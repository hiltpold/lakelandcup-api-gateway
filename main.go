package main

import (
	"github.com/hiltpold/lakelandcup-api-gateway/commands"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := commands.RootCommand().Execute(); err != nil {
		logrus.Fatal(err)
	}
	/*
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
	*/
}
