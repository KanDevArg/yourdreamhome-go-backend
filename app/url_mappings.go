package app

import (
	"github.com/KanDevArg/yourdreamhome/go-backend/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initMiddlewares() {
	router.Use(VendorMiddleware())
	router.Use(RequestIdMiddleware())
}

func mapUrls() {
	router.GET("/listings/:postal_code", controllers.GetListing)
	router.GET("/listings", controllers.GetAllListings)
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Pong")
	})
}
