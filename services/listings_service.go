package services

import "github.com/KanDevArg/yourdreamhome/go-backend/domain"

func GetListing(postalCode string) (*domain.Listing, error) {
	return domain.GetListingsByPostalCode(postalCode)
}
func GetAllListings() []domain.Listing {
	return domain.GetAllListings()
}
