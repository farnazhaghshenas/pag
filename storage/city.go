package storage

import "pack-and-go/models"

func (sql *SqlClient) CreateCity(city models.City) (int64, error) {
	result, err := sql.Client.Exec("INSERT `cities` (`name`) VALUES (?)", city.Name)
	if err != nil {
		return 0, err
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return LastInsertId, nil
}
