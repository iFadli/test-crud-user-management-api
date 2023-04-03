package database

import (
	"database/sql"
	"fmt"
	"log"
	"user-management-api/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME))
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil, err
	}

	return db, nil
}
