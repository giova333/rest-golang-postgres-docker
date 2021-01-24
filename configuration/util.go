package configuration

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	Configuration struct {
		ServerPort, DbHost, DbUser, DbPassword, DbName string
	}
)

func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

func InitConfig() *Configuration {
	file, err := os.Open("application-properties.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	var config Configuration
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("[logAppConfig]: %s\n", err)
	}

	loadConfigFromEnvironment(&config)
	return &config
}

func loadConfigFromEnvironment(config *Configuration) {
	serverPort, ok := os.LookupEnv("SERVER_PORT")
	if ok {
		config.ServerPort = serverPort
		log.Printf("[INFO]: ServerPort information loaded from env.")
	}
	dbHost, ok := os.LookupEnv("DB_HOST")
	if ok {
		config.DbHost = dbHost
		log.Printf("[INFO]: DbHost information loaded from env.")
	}
	dbUser, ok := os.LookupEnv("DB_USER")
	if ok {
		config.DbUser = dbUser
		log.Printf("[INFO]: DbUser information loaded from env.")
	}
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if ok {
		config.DbPassword = dbPassword
		log.Printf("[INFO]: DbPassword information loaded from env.")
	}
	dbName, ok := os.LookupEnv("DB_NAME")
	if ok {
		config.DbName = dbName
		log.Printf("[INFO]: DbName information loaded from env.")
	}
}
