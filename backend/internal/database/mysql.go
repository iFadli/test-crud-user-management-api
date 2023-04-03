package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"
	"user-management-api/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	var db *sql.DB
	var err error

	maxRetry := 10
	for i := 0; i < maxRetry; i++ {
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME))

		if i > 0 {
			fmt.Println("DB Connection : Retry Mechanism [" + strconv.Itoa(i) + "x]")
		}
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil, err
	}

	return db, nil
}
