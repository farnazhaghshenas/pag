package http

import (
	"github.com/labstack/echo/v4"
	"pack-and-go/storage"
)

// API is structure that implement API interface
type API struct {
	Server *echo.Echo
	DB     storage.Provider
}

type CreateTripRequest struct {
	OriginId      int     `json:"originId" validate:"required"`
	DestinationId int     `json:"destinationId" validate:"required"`
	Dates         string  `json:"dates" validate:"required"`
	Price         float64 `json:"price" validate:"required"`
}