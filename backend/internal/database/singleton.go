package database

import (
	"database/sql"
	"sync"
)

var db *sql.DB
var once sync.Once

func GetInstance() (*sql.DB, error) {
	once.Do(func() {
		var err error
		db, err = Connect()
		if err != nil {
			panic(err)
		}
	})

	return db, nil
}
