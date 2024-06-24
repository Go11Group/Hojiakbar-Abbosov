package main

import (
	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/users_service/api"
	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/users_service/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	server := api.Routes(db)
	err = server.ListenAndServe() // 80
	if err != nil {
		panic(err)
	}
}
