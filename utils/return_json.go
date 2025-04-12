package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnSuccess(ctx *gin.Context, msg string, data interface{}) {
	var defaultMsg = "success"
	if msg != "" {
		defaultMsg = msg
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  defaultMsg,
		"data": data,
	})
}
func ReturnError(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  msg,
	})
}
