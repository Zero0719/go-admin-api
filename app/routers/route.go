package routers

import (
	"go-admin-api/app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(server *gin.Engine)  {
	server.GET("/", (&controllers.IndexController{}).Index)
}