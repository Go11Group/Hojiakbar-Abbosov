package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"go-gin-crud/config"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", config.GetDBURI())
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	log.Println("Connected to the database")
}

func Close() {
	DB.Close()
}
