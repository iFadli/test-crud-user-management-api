package database

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnect(t *testing.T) {
	// Set test cases
	testCases := []struct {
		name        string
		mockFunc    func() (*sql.DB, error)
		expectedErr error
	}{
		{
			name: "success",
			mockFunc: func() (*sql.DB, error) {
				db, mock, err := sqlmock.New()
				if err != nil {
					return nil, err
				}
				defer db.Close()

				mock.ExpectPing()
				return db, nil
			},
			expectedErr: nil,
		},
		{
			name: "connection error",
			mockFunc: func() (*sql.DB, error) {
				db, mock, err := sqlmock.New()
				if err != nil {
					return nil, err
				}
				defer db.Close()

				mock.ExpectPing().WillReturnError(errors.New("database connection error"))
				return db, nil
			},
			expectedErr: errors.New("database connection error"),
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call Connect function
			_, err := Connect()
			if err != tc.expectedErr {
				t.Errorf("Expected %v, but got %v", tc.expectedErr, err)
			}
		})
	}
}
