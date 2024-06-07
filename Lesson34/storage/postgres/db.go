package postgres

import (
    "database/sql"
    _ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(connString string) error {
    var err error
    db, err = sql.Open("postgres", connString)
    if err != nil {
        return err
    }
    if err = db.Ping(); err != nil {
        return err
    }
    return nil
}

func GetDB() *sql.DB {
    return db
}

func CloseDB() {
    if db != nil {
        db.Close()
    }
}
