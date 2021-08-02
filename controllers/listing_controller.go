package controllers

import (
	"encoding/json"
	"github.com/KanDevArg/yourdreamhome/go-backend/services"
	"github.com/KanDevArg/yourdreamhome/go-backend/utils"
	"net/http"
)

func GetListings(res http.ResponseWriter, req *http.Request) {
	pc := req.URL.Query().Get("postal_code")

	listing, err := services.GetListings(pc)
	if err != nil {
		errData := utils.ApplicationError{
			Message:    err.Error(),
			StatusCode: http.StatusNotFound,
		}
		res.WriteHeader(http.StatusNotFound)
		errorObject, _ := json.Marshal(errData)
		res.Write(errorObject)
		return
	}

	jsonValue, _ := json.Marshal(listing)
	res.Write(jsonValue)
}

func GetAllListings(res http.ResponseWriter, req *http.Request) {
	listings := services.GetAllListings()
	jsonValue, _ := json.Marshal(listings)
	res.Write(jsonValue)
}
