package domain

import (
	"fmt"
	"github.com/KanDevArg/yourdreamhome/go-backend/utils"
	"sort"
)

var (
	listings = map[string]*Listing{
		"5007": {Id: 1, Address: "Cesar Carrizo 2982", City: "Cordoba", PostCode: "5007"},
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
	var keys []string

	for k := range listings {
		keys = append(keys, k)
	}

	//sort.Sort(sort.Reverse(sort.Strings(keys))
	sort.Strings(keys)

	for _, k := range keys {
		arr = append(arr, *listings[k])
	}
	return arr
}

type Listing struct {
	Id       int
	Address  string
	City     string
	PostCode string
}
