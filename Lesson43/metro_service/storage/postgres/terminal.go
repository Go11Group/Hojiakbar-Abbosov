package postgres

import (
	"database/sql"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/metro_service/models"
	"github.com/google/uuid"
)

type TerminalRepo struct {
	Db *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{Db: db}
}

func (t *TerminalRepo) Create(terminal *models.CreateTerminal) error {
	_, err := t.Db.Exec("INSERT INTO terminal(id, station_id) VALUES ($1, $2)",
		uuid.NewString(), terminal.StationId)
	return err
}

func (t *TerminalRepo) GetById(id string) (*models.Terminal, error) {
	var terminal models.Terminal
	err := t.Db.QueryRow("SELECT id, station_id FROM terminal WHERE id = $1", id).
		Scan(&terminal.Id, &terminal.StationId)
	if err != nil {
		return nil, err
	}
	return &terminal, nil
}

func (t *TerminalRepo) Update(terminal *models.Terminal) error {
	_, err := t.Db.Exec("UPDATE terminal SET station_id = $1 WHERE id = $2",
		terminal.StationId, terminal.Id)
	return err
}

func (t *TerminalRepo) Delete(id string) error {
	_, err := t.Db.Exec("DELETE FROM terminal WHERE id = $1", id)
	return err
}
