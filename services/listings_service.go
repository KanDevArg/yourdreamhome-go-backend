package services

import "github.com/KanDevArg/yourdreamhome/go-backend/domain"

func GetListings(postalCode string) (*domain.Listing, error) {
	return domain.GetListingsByPostalCode(postalCode)
}
