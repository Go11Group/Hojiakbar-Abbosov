package database

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("postgres", "user=postgres password=root dbname=exam sslmode=disable")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
}
