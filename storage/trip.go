package storage

import (
	mainSql "database/sql"
	"pack-and-go/models"
)

func (sql *SqlClient) CreateTrip(trip models.Trip) (int64, error) {
	result, err := sql.Client.Exec("INSERT `trips` (`originId`,`destinationId`, `dates`, `price`) VALUES (?, ?, ?, ?)",
		trip.OriginId, trip.DestinationId, trip.Dates, trip.Price)
	
	if err != nil {
		return 0, err
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return LastInsertId, nil
}

func (sql *SqlClient) ReadTrip(id int64) (models.GetTrip, error) {
	row := sql.Client.QueryRow("SELECT trips.id as id, co.name as origin, cd.name as destination, trips.dates, trips.price FROM trips JOIN cities co ON co.id = trips.originId JOIN cities cd ON cd.id = trips.destinationId WHERE trips.id = ? ", id)
	
	var trip models.GetTrip
	err := row.Scan(&trip.Id, &trip.Origin, &trip.Destination, &trip.Dates, &trip.Price)
	if err != nil {
		if err == mainSql.ErrNoRows {
			return models.GetTrip{}, nil
		}
		return models.GetTrip{}, err
	}
	return trip, nil
}

func (sql *SqlClient) ReadAllTrips() ([]models.GetTrip, error) {
	rows, err := sql.Client.Query("SELECT trips.id as id, co.name as origin, cd.name as destination, trips.dates, trips.price FROM trips JOIN cities co ON co.id = trips.originId JOIN cities cd ON cd.id = trips.destinationId")
	if err != nil {
		return []models.GetTrip{}, err
	}

	var trips []models.GetTrip
	for rows.Next() {
		var trip models.GetTrip
		err := rows.Scan(&trip.Id, &trip.Origin, &trip.Destination, &trip.Dates, &trip.Price)
		if err != nil {
			return []models.GetTrip{}, err
		}
		trips = append(trips, trip)
	}

	return trips, nil
}