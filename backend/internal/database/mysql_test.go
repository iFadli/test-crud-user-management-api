package database

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnect(t *testing.T) {
	// set test cases
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
	}

	// run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// set mock sql.Open function
			//sqlOpen := tc.mockFunc

			// call Connect function
			_, err := MockConnect()
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, but got %v", tc.expectedErr, err)
			}
		})
	}
}
