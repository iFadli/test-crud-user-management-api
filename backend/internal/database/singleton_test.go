package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetInstance(t *testing.T) {
	// create mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock db: %v", err)
	}
	defer db.Close()

	// call GetInstance function
	actualDB, err := MockGetInstance(db)
	if err != nil {
		t.Fatalf("error getting database instance: %v", err)
	}

	// check that the actual DB instance returned by GetInstance matches the mock DB instance
	if actualDB != db {
		t.Errorf("expected DB instance %v, but got %v", db, actualDB)
	}

	// check that expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("failed to meet expectations: %v", err)
	}
}
