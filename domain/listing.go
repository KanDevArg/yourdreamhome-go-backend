package domain

import (
	"fmt"
	"github.com/KanDevArg/yourdreamhome/go-backend/utils"
)

var (
	listings = map[string]*Listing{
		"5001": {Id: 1, Address: "Cesar Carrizo 2982", City: "Cordoba", PostCode: "5001"},
		"5002": {Id: 2, Address: "Loka loka 2001", City: "Cordoba", PostCode: "5002"},
		"5003": {Id: 3, Address: "Colodrero 6001", City: "Cordoba", PostCode: "5003"},
	}
)

func GetListingsByPostalCode(postalCode string) (*Listing, error) {
	l := listings[postalCode]
	if l == nil {
		return nil,
			utils.ApplicationError{
				Message:    fmt.Sprintf("no listing found on postalcode %v", postalCode),
				StatusCode: 100,
			}
	}
	return l, nil
}

func GetAllListings() []Listing {
	var arr []Listing
	for _, v := range listings {
		arr = append(arr, *v)
	}
	return arr
}

type Listing struct {
	Id       int
	Address  string
	City     string
	PostCode string
}
