package controllers

import (
	"go-admin-api/app/utils/response"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (*IndexController) Index(c *gin.Context)  {
	response.Success(c, "Hello, World1111222!")
}