package handler

import (
	"database/sql"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/users_service/storage/postgres"
)

type handler struct {
	User *postgres.UserRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		User: postgres.NewUserRepo(db),
	}
}
