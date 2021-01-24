package routers

import (
	"github.com/gorilla/mux"
	"rest/repository"
)

func InitRoutes(repo *repository.UserRepository) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = SetUsersRouters(router, repo)
	return router
}
