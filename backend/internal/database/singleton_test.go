package database

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetInstance(t *testing.T) {
	// Setup mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Set expectations
	mock.ExpectPing()

	// Call GetInstance function
	actualDB, err := GetInstance()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Check if returned database is correct
	if actualDB != db {
		t.Errorf("Expected %v, but got %v", db, actualDB)
	}

	// Check if expected queries were executed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}

	// Test case ketika terjadi error saat koneksi ke database
	expectedErr := errors.New("database connection error")
	mock.ExpectPing().WillReturnError(expectedErr)
	actualDB, err = GetInstance()
	if actualDB != nil {
		t.Errorf("Expected nil database, but got %v", actualDB)
	}
	if err != expectedErr {
		t.Errorf("Expected error %v, but got %v", expectedErr, err)
	}
}
