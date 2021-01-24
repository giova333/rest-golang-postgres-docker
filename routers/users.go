package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest/controllers"
	"rest/repository"
)

func SetUsersRouters(router *mux.Router, repo *repository.UserRepository) *mux.Router {
	router.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		controllers.GetUsers(writer, repo)
	}).Methods("GET")

	router.HandleFunc("/users/{uuid}", func(writer http.ResponseWriter, request *http.Request) {
		controllers.GetUser(writer, request, repo)
	}).Methods("GET")

	router.HandleFunc("/users/{uuid}", func(writer http.ResponseWriter, request *http.Request) {
		controllers.DeleteUser(writer, request, repo)
	}).Methods("DELETE")

	router.HandleFunc("/users", func(writer http.ResponseWriter, request *http.Request) {
		controllers.CreateUsers(writer, request, repo)
	}).Methods("POST")
	return router
}
