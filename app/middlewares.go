package app

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("X-Request-Id", uuid.NewV4().String())
		ctx.Next()
	}
}

func VendorMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
