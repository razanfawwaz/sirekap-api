package main

import (
	"github.com/gin-gonic/gin"
	"github.com/razanfawwaz/sirekap-api/config"
	"github.com/razanfawwaz/sirekap-api/interfaces"
)

func main() {

	mode := config.GetConfig("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Group
	routeGroup := r.Group("/api/v1")
	routeGroup.GET("/ppwp", interfaces.GetAllData)

	// connect to database
	config.ConnectDB()
	r.Run(config.GetConfig("PORT"))
}
