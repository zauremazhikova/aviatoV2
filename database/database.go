package database

import (
	"aviatoV2/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func DB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}
	return db
}
