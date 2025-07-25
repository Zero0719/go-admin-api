package app

import (
	"go-admin-api/app/config"
	"go-admin-api/app/routers"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func Bootstrap() {
	initConfig() // 初始化配置
	initServer() // 初始化服务
	initRouter() // 初始化路由
	startServer() // 启动服务
}

func initConfig() {
	config.Init()
}

func initServer()  {
	server = gin.New()
}

func initRouter() {
	routers.SetupRouter(server)
}

func startServer() {
	server.Run(":" + config.Cfg.Server.Port)
}