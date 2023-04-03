package database

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestPrepareDB(t *testing.T) {
	// Setup mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Set expectations
	mock.ExpectPing()
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS `user`").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT COUNT").WillReturnRows(sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(1))

	// Call PrepareDB function
	err = PrepareDB(db)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Check if expected queries were executed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}

	// Test case ketika tabel user gagal dibuat
	mock.ExpectPing()
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS `user`").WillReturnError(errors.New("create table error"))
	err = PrepareDB(db)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}

	// Test case ketika query SELECT COUNT(*) gagal
	mock.ExpectPing()
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS `user`").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectQuery("SELECT COUNT").WillReturnError(errors.New("query error"))
	err = PrepareDB(db)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}

func TestPrepareDefaultUser(t *testing.T) {
	// Setup mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Test case ketika terjadi error saat hashing password
	mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(0, 1))
	err = PrepareDefaultUser(db)
	if err == nil {
		t.Errorf("Expected an error, but got none")
	}
}
