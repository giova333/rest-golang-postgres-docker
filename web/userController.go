package web

import (
	"encoding/json"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"rest/configuration"
	"rest/domain"
	"rest/persistence"
)

type UserController struct {
	repository *persistence.UserRepository
}

func NewUserController(repo *persistence.UserRepository) *UserController {
	return &UserController{repository: repo}
}

func (controller *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := controller.repository.GetAll()
	if err != nil {
		configuration.DisplayAppError(w, err, err.Error(), 500)
		return
	}
	j, err := json.Marshal(users)
	if err != nil {
		configuration.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func (controller *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	user, err := controller.repository.GetOne(uuid)
	if err != nil {
		configuration.DisplayAppError(w, err, err.Error(), 404)
		return
	}
	j, err := json.Marshal(user)
	if err != nil {
		configuration.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func (controller *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	err := controller.repository.Delete(uuid)
	if err != nil {
		configuration.DisplayAppError(w, err, err.Error(), 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (controller *UserController) CreateUsers(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		configuration.DisplayAppError(w, err, "Invalid User data", 500)
		return
	}

	user.UUID = guuid.New()
	err = controller.repository.Create(user)

	if err != nil {
		configuration.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		configuration.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
