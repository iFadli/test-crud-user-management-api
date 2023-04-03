package controllers

import (
	"errors"
	"user-management-api/internal/utils"

	"user-management-api/internal/models"
	"user-management-api/internal/repositories"
)

type UserController struct {
	userRepo *repositories.UserRepository
}

func NewUserController() *UserController {
	return &UserController{
		userRepo: repositories.NewUserRepository(),
	}
}

func (c *UserController) GetAllUsers() ([]*models.User, error) {
	return c.userRepo.GetAllUsers()
}

func (c *UserController) GetUserByID(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("Invalid user ID")
	}

	return c.userRepo.GetUserByID(id)
}

func (c *UserController) SaveUser(user *models.User) error {
	if user == nil || user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("Invalid user data")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := c.userRepo.SaveUser(user); err != nil {
		return err
	}

	return nil
}

func (c *UserController) UpdateUser(id int, user *models.User) error {
	if id <= 0 || user == nil || user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("Invalid user data")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	return c.userRepo.UpdateUser(id, user)
}

func (c *UserController) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("Invalid user ID")
	}

	return c.userRepo.DeleteUser(id)
}

func (c *UserController) Login(auth *models.User) (string, error) {
	if auth == nil || auth.Username == "" || auth.Password == "" {
		return "", errors.New("Invalid auth data")
	}

	storedAuth, err := c.userRepo.GetUserByUsername(auth.Username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if !utils.ComparePasswordAndHash(auth.Password, storedAuth.Password) {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateToken(storedAuth)
	if err != nil {
		return "", errors.New("failed generate token, please try again later")
	}

	return token, nil
}
