package storage

import "pack-and-go/models"

type Provider interface {
	CreateTrip(models.Trip) (int64, error)
	ReadTrip(id int64) (models.GetTrip, error)
	ReadAllTrips() ([]models.GetTrip, error)

	CreateCity(city models.City) (int64, error)
}
