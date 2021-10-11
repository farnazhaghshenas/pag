package storage

import (
	"database/sql"
	"fmt"
	"pack-and-go/config"

	_ "github.com/go-sql-driver/mysql"
)

type SqlClient struct {
	Client *sql.DB
}

// NewDBConnection is the init SQL  factory method.
func NewDBConnection(config config.DB) (*SqlClient, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Name))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	sqlClient := &SqlClient{Client: db}

	return sqlClient, err
}


