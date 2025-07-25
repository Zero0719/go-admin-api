package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Code: 0,
		Msg: "ok",
		Data: data,
	})
}

func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, ErrorResponse{
		Code: code,
		Msg: msg,
	})
}