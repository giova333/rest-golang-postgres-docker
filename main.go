package main

import (
	"log"
	"net/http"
	"rest/common"
	"rest/repository"
	"rest/routers"
)

func main() {
	repo := &repository.UserRepository{}
	router := routers.InitRoutes(repo)
	config := common.InitConfig()

	server := &http.Server{
		Addr:    config.Server,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
