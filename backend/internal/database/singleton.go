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

func MockConnect() (*sql.DB, error) {
	// code to connect to database
	return nil, nil
}

func MockGetInstance(db *sql.DB) (*sql.DB, error) {
	once.Do(func() {
		if db == nil {
			var err error
			db, err = MockConnect()
			if err != nil {
				panic(err)
			}
		}
	})

	return db, nil
}
