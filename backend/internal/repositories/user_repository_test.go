package repositories_test

import (
	"errors"
	"testing"
	"user-management-api/internal/models"
	"user-management-api/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	user1 := &models.User{
		ID:        1,
		Username:  "user1",
		Email:     "user1@example.com",
		CreatedAt: "2022-04-01 10:00:00",
		UpdatedAt: "2022-04-01 10:00:00",
	}

	user2 := &models.User{
		ID:        2,
		Username:  "user2",
		Email:     "user2@example.com",
		CreatedAt: "2022-04-02 10:00:00",
		UpdatedAt: "2022-04-02 10:00:00",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "created_at", "updated_at"}).
		AddRow(user1.ID, user1.Username, user1.Email, user1.CreatedAt, user1.UpdatedAt).
		AddRow(user2.ID, user2.Username, user2.Email, user2.CreatedAt, user2.UpdatedAt)

	mock.ExpectQuery("SELECT id, username, email, created_at, updated_at FROM user").
		WillReturnRows(rows)

	repo := &repositories.UserRepository{
		Db: db,
	}

	users, err := repo.GetAllUsers()

	assert.NoError(t, err)
	assert.NotNil(t, users)
	assert.Len(t, users, 2)
	assert.Equal(t, user1, users[0])
	assert.Equal(t, user2, users[1])

	// test error case
	mock.ExpectQuery("SELECT id, username, email, created_at, updated_at FROM user").
		WillReturnError(errors.New("database error"))

	users, err = repo.GetAllUsers()

	assert.Error(t, err)
	assert.Nil(t, users)
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	user2 := &models.User{
		ID:        1,
		Username:  "user2",
		Email:     "user2@example.com",
		CreatedAt: "2022-04-02 10:00:00",
		UpdatedAt: "2022-04-02 10:00:00",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "created_at", "updated_at"}).
		AddRow(user2.ID, user2.Username, user2.Email, user2.CreatedAt, user2.UpdatedAt)

	mock.ExpectQuery("SELECT id, username, email, created_at, updated_at FROM user WHERE id = 1").
		WillReturnRows(rows)

	repo := &repositories.UserRepository{
		Db: db,
	}

	user, err := repo.GetUserByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user2, user)

	// test error case
	mock.ExpectQuery("SELECT id, username, email, created_at, updated_at FROM user WHERE id = 2").
		WillReturnError(errors.New("database error"))

	user, err = repo.GetUserByID(2)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestGetUserByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	user2 := &models.User{
		ID:       1,
		Username: "user2",
		Email:    "user2@example.com",
		Password: "1q2w3e4r5t",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "email", "password"}).
		AddRow(user2.ID, user2.Username, user2.Email, user2.Password)

	mock.ExpectQuery("SELECT id, username, email, password FROM user WHERE username = 'user2'").
		WillReturnRows(rows)

	repo := &repositories.UserRepository{
		Db: db,
	}

	user, err := repo.GetUserByUsername("user2")

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user2, user)

	// test error case
	mock.ExpectQuery("SELECT id, username, email, password FROM user WHERE username = 'hey'").
		WillReturnError(errors.New("database error"))

	user, err = repo.GetUserByUsername("hey")

	assert.Error(t, err)
	assert.Nil(t, user)
}
