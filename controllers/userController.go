package controllers

import (
	"encoding/json"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"rest/common"
	"rest/models"
	"rest/repository"
)

func GetUsers(w http.ResponseWriter, repo *repository.UserRepository) {
	users := repo.GetAll()
	j, err := json.Marshal(users)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetUser(w http.ResponseWriter, r *http.Request, repo *repository.UserRepository) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	user, err := repo.GetOne(uuid)
	if err != nil {
		common.DisplayAppError(w, err, err.Error(), 404)
		return
	}
	j, err := json.Marshal(user)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func DeleteUser(w http.ResponseWriter, r *http.Request, repo *repository.UserRepository) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	err := repo.Delete(uuid)
	if err != nil {
		common.DisplayAppError(w, err, err.Error(), 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func CreateUsers(w http.ResponseWriter, r *http.Request, repo *repository.UserRepository) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid User data", 500)
		return
	}

	user.UUID = guuid.New()
	repo.Create(user)

	j, err := json.Marshal(user)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
