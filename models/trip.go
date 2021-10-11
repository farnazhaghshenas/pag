package models

type Trip struct {
	Id            int32   `json:"id" sql:"id"`
	OriginId      int32   `json:"originId" sql:"originId"`
	DestinationId int32   `json:"destinationId" sql:"destinationId"`
	Dates         string  `json:"dates" sql:"dates"`
	Price         float64 `json:"price" sql:"price"`
}

type GetTrip struct {
	Id          int32   `json:"id" sql:"id"`
	Origin      string  `json:"origin" sql:"origin"`
	Destination string  `json:"destination" sql:"destination"`
	Dates       string  `json:"dates" sql:"dates"`
	Price       float64 `json:"price" sql:"price"`
}