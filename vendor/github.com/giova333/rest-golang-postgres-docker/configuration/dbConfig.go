package configuration

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitDatabase(config *Configuration) *sql.DB {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbUser, config.DbPassword, config.DbName)

	connection, err := createConnection(connectionString)

	if err != nil {
		log.Fatalf("[creatingDbConnection]: %s\n", err)
	}
	return connection
}

func createConnection(desc string) (*sql.DB, error) {
	db, err := sql.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db, nil
}
