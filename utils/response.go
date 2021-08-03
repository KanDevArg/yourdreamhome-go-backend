package utils

import "github.com/gin-gonic/gin"

func DoResponse(ctx *gin.Context, status int, data interface{}) {
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(status, data)
		return
	}
	ctx.JSON(status, data)
}
