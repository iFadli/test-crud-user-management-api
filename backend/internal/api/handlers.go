package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"user-management-api/internal/controllers"
	"user-management-api/internal/models"

	"github.com/gorilla/mux"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func WriteResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := APIResponse{
		Status: status,

		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var authRequest AuthRequest

	err := json.NewDecoder(r.Body).Decode(&authRequest)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	userController := controllers.NewUserController()

	token, err := userController.Login(&models.User{
		Username: authRequest.Username,
		Password: authRequest.Password,
	})

	if err != nil {
		WriteResponse(w, http.StatusUnauthorized, "Failed to login. "+err.Error(), nil)
		return
	}

	WriteResponse(w, http.StatusOK, "Login success", map[string]string{"token": token})
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	userController := controllers.NewUserController()

	users, err := userController.GetAllUsers()
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, "Failed to get users data", nil)
		return
	}

	WriteResponse(w, http.StatusOK, "Get all users success", users)
}

func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid id parameter", nil)
		return
	}

	userController := controllers.NewUserController()

	user, err := userController.GetUserByID(id)
	if err != nil {
		WriteResponse(w, http.StatusNotFound, "User not found", nil)
		return
	}

	WriteResponse(w, http.StatusOK, "Get user by ID success", user)
}

func SaveUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	userController := controllers.NewUserController()

	if err := userController.SaveUser(&user); err != nil {
		WriteResponse(w, http.StatusInternalServerError, "Failed to save user", nil)
		return
	}

	WriteResponse(w, http.StatusCreated, "Save user success", nil)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid id parameter", nil)
		return
	}

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	userController := controllers.NewUserController()

	if err := userController.UpdateUser(id, &user); err != nil {
		WriteResponse(w, http.StatusInternalServerError, "Failed to update user", nil)
		return
	}

	WriteResponse(w, http.StatusOK, "Update user success", nil)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Invalid id parameter", nil)
		return
	}

	if id == 1 {
		WriteResponse(w, http.StatusForbidden, "Permission Denied for deleting Master User", nil)
		return
	}

	userController := controllers.NewUserController()

	if err := userController.DeleteUser(id); err != nil {
		WriteResponse(w, http.StatusInternalServerError, "Failed to delete user", nil)
		return
	}

	WriteResponse(w, http.StatusOK, "Delete user success", nil)
}
