package main

import (
	"log"
	"net/http"
	"user-management-api/internal/api"
	"user-management-api/internal/database"
	"user-management-api/internal/middleware"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.GetInstance()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := database.PrepareDB(db); err != nil {
		log.Fatal(err)
	}

	masterRouter := mux.NewRouter()
	api.GuestRoutes(masterRouter)
	securedRouter := mux.NewRouter()
	api.SecureRoutes(securedRouter)

	// Gunakan router utama sebagai parent dari securedRouter
	masterRouter.PathPrefix("/").Handler(middleware.SecurityMiddleware(securedRouter))

	log.Fatal(http.ListenAndServe(":8080", masterRouter))
}
