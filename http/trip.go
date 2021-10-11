package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"pack-and-go/models"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func (api *API) CreateTrip(ctx echo.Context) error {
	request := new(CreateTripRequest)
	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	c := models.Trip{
		OriginId: int32(request.OriginId),
		DestinationId: int32(request.DestinationId),
		Dates: request.Dates,
		Price: request.Price,
	}

	_ , err = api.DB.CreateTrip(c)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, nil)
}

func (api *API) GetTrips(ctx echo.Context) error {
	trips , err := api.DB.ReadAllTrips()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, trips)
}

func (api *API) GetTrip(ctx echo.Context) error {
	id := ctx.Param("id")
	tripId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	trip , err := api.DB.ReadTrip(tripId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, trip)
}

