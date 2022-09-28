package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hiltpold/lakelandcup-api-gateway/config"
	"github.com/hiltpold/lakelandcup-api-gateway/services/auth"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
