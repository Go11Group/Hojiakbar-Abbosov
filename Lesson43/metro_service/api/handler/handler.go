package handler

import (
	"database/sql"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/metro_service/storage/postgres"
)

type handler struct {
	Station *postgres.StationRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		Station: postgres.NewStationRepo(db),
	}
}

type CardHandler struct {
	Repo *postgres.CardRepo
}

func NewCardHandler(db *sql.DB) *CardHandler {
	return &CardHandler{
		Repo: postgres.NewCardRepo(db),
	}
}

type TransactionHandler struct {
	Repo *postgres.TransactionRepo
}

func NewTransactionHandler(db *sql.DB) *TransactionHandler {
	return &TransactionHandler{
		Repo: postgres.NewTransactionRepo(db),
	}
}

type TerminalHandler struct {
	Repo *postgres.TerminalRepo
}

func NewTerminalHandler(db *sql.DB) *TerminalHandler {
	return &TerminalHandler{
		Repo: postgres.NewTerminalRepo(db),
	}
}