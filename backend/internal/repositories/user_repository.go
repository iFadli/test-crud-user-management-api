package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"user-management-api/internal/database"
	"user-management-api/internal/models"
)

// Buat interface baru dengan method-method yang dibutuhkan oleh UserRepository
type DBInterface interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type UserRepository struct {
	Db DBInterface
}

func NewUserRepository() *UserRepository {
	db, err := database.GetInstance()
	if err != nil {
		panic(err)
	}

	return &UserRepository{
		Db: db,
	}
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	rows, err := r.Db.Query("SELECT id, username, email, created_at, updated_at FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}

	row := r.Db.QueryRow(fmt.Sprintf("SELECT id, username, email, created_at, updated_at FROM user WHERE id = %d", id))
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no data found")
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}

	row := r.Db.QueryRow(fmt.Sprintf("SELECT id, username, email, password FROM user WHERE username = '%s'", username))
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no data found")
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (r *UserRepository) SaveUser(user *models.User) error {
	query := fmt.Sprintf("INSERT INTO user (username, email, password) VALUES ('%s', '%s', '%s')", user.Username, user.Email, user.Password)

	_, err := r.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(id int, user *models.User) error {
	query := fmt.Sprintf("UPDATE user SET username = '%s', email = '%s', password = '%s', updated_at = NOW() WHERE id = %d", user.Username, user.Email, user.Password, id)

	result, err := r.Db.Exec(query)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("no data updated")
	}

	return nil
}

func (r *UserRepository) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM user WHERE id = %d", id)

	result, err := r.Db.Exec(query)
	if err != nil {
		return err
	}

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {
		return errors.New("no data deleted")
	}

	return nil
}
