package api

import (
	"net/http"
	"user-management-api/internal/middleware"

	"github.com/gorilla/mux"
)

func GuestRoutes(router *mux.Router) {
	router.HandleFunc("/login", LoginHandler).Methods(http.MethodPost)
}

func SecureRoutes(router *mux.Router) {
	router.Use(middleware.ValidateToken)
	router.HandleFunc("/users", GetAllUsersHandler).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", GetUserByIDHandler).Methods(http.MethodGet)
	router.HandleFunc("/users", SaveUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", DeleteUserHandler).Methods(http.MethodDelete)
}
