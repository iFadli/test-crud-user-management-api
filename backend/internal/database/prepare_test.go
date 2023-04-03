package database

import (
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestPrepareDB(t *testing.T) {
	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	// expect Ping function call
	mock.ExpectPing()

	// expect CreateTable function call
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS `user` ( `id` int NOT NULL AUTO_INCREMENT, `username` varchar(50) NOT NULL, `email` varchar(100) NOT NULL, `password` varchar(255) NOT NULL, `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP, `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (`id`), UNIQUE KEY `username_UNIQUE` (`username`), UNIQUE KEY `email_UNIQUE` (`email`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// expect QueryRow function call
	mock.ExpectQuery("SELECT COUNT(id) FROM user").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	// call PrepareDB function
	err = PrepareDB(db)
	if err != nil {
		t.Fatalf("error preparing database: %v", err)
	}

	// check that expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("failed to meet expectations: %v", err)
	}
}

func TestPrepareDefaultUser(t *testing.T) {
	// create mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	// expect Insert function call
	mock.ExpectExec("INSERT INTO user (.*)").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// call PrepareDefaultUser function
	err = PrepareDefaultUser(db)
	if err != nil {
		t.Fatalf("error preparing default user: %v", err)
	}

	// check that expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("failed to meet expectations: %v", err)
	}
}
