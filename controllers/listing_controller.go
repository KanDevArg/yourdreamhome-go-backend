package controllers

import (
	"github.com/KanDevArg/yourdreamhome/go-backend/services"
	"github.com/KanDevArg/yourdreamhome/go-backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetListing(ctx *gin.Context) {
	pc := ctx.Param("postal_code")

	listing, err := services.GetListing(pc)
	if err != nil {
		errData := utils.ApplicationError{
			Message:    err.Error(),
			StatusCode: http.StatusNotFound,
		}

		utils.DoResponse(ctx, http.StatusNotFound, errData)
		return
	}
	utils.DoResponse(ctx, http.StatusOK, listing)
}

func GetAllListings(ctx *gin.Context) {
	listings := services.GetAllListings()
	utils.DoResponse(ctx, http.StatusOK, listings)
}
