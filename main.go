package main

import (
	"github.com/giova333/rest-golang-postgres-docker/configuration"
	"github.com/giova333/rest-golang-postgres-docker/persistence"
	"github.com/giova333/rest-golang-postgres-docker/web"
	"log"
	"net/http"
)

func main() {
	config := configuration.InitConfig()
	db := configuration.InitDatabase(config)
	repo := persistence.NewUserRepository(db)
	controller := web.NewUserController(repo)
	router := web.RegisterRoutes(controller)

	server := &http.Server{
		Addr:    config.ServerPort,
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
