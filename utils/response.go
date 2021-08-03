package utils

import "github.com/gin-gonic/gin"

func DoResponse(ctx *gin.Context, status int, data interface{}) {
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(status, data)
		return
	}
	ctx.JSON(status, data)
}

func DoErrorResponse(ctx *gin.Context, err ApplicationError) {
	if ctx.GetHeader("Accept") == "application/xml" {
		ctx.XML(err.Status, err.Error())
		return
	}
	ctx.JSON(err.Status, err.Error())
}
