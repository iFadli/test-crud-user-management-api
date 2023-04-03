package api

import (
	"net/http"
	"testing"
)

func TestDeleteUserHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteUserHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestGetAllUsersHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAllUsersHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestGetUserByIDHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetUserByIDHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestLoginHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoginHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestSaveUserHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveUserHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestUpdateUserHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateUserHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestWriteResponse(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		status  int
		message string
		data    interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteResponse(tt.args.w, tt.args.status, tt.args.message, tt.args.data)
		})
	}
}
