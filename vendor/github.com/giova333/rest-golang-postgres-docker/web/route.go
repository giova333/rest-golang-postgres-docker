package web

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(controller *UserController) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/{uuid}", controller.GetUser).Methods("GET")
	router.HandleFunc("/users", controller.CreateUsers).Methods("POST")
	router.HandleFunc("/users/{uuid}", controller.DeleteUser).Methods("DELETE")
	return router
}
