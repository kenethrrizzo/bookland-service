package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/kenethrrizzo/bookland-service/cmd/api/config"
)

func Connect(config *config.Datasource) (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User, config.Password, config.Host, config.Port, config.Database)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
