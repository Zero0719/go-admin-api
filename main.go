package main

import (
	"go-admin-api/app/config"
	"go-admin-api/app/utils/response"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		response.Success(c, "Hello, World!")
	})
	server.Run(":" + config.Cfg.Server.Port)
}
