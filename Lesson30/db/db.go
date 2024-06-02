package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
    connStr := "user=postgres password=root dbname=mydb sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    return db, nil
}
