package controllers

import (
	"reflect"
	"testing"
	"user-management-api/internal/models"
	"user-management-api/internal/repositories"
)

func TestNewUserController(t *testing.T) {
	tests := []struct {
		name string
		want *UserController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserController(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_DeleteUser(t *testing.T) {
	type fields struct {
		userRepo *repositories.UserRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UserController{
				userRepo: tt.fields.userRepo,
			}
			if err := c.DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserController_GetAllUsers(t *testing.T) {
	type fields struct {
		userRepo *repositories.UserRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UserController{
				userRepo: tt.fields.userRepo,
			}
			got, err := c.GetAllUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUsers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_GetUserByID(t *testing.T) {
	type fields struct {
		userRepo *repositories.UserRepository
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UserController{
				userRepo: tt.fields.userRepo,
			}
			got, err := c.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_Login(t *testing.T) {
	type fields struct {
		userRepo *repositories.UserRepository
	}
	type args struct {
		auth *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UserController{
				userRepo: tt.fields.userRepo,
			}
			got, err := c.Login(tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_SaveUser(t *testing.T) {
	type fields struct {
		userRepo *repositories.UserRepository
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UserController{
				userRepo: tt.fields.userRepo,
			}
			if err := c.SaveUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("SaveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserController_UpdateUser(t *testing.T) {
	type fields struct {
		userRepo *repositories.UserRepository
	}
	type args struct {
		id   int
		user *models.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UserController{
				userRepo: tt.fields.userRepo,
			}
			if err := c.UpdateUser(tt.args.id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
