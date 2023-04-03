package api

import (
	"github.com/gorilla/mux"
	"testing"
)

func TestGuestRoutes(t *testing.T) {
	type args struct {
		router *mux.Router
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GuestRoutes(tt.args.router)
		})
	}
}

func TestSecureRoutes(t *testing.T) {
	type args struct {
		router *mux.Router
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SecureRoutes(tt.args.router)
		})
	}
}
