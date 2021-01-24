package main

import (
	"log"
	"net/http"
	"rest/configuration"
	"rest/persistence"
	"rest/web"
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
