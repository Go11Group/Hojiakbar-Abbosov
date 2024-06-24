package postgres

import (
	"database/sql"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/metro_service/models"
	"github.com/google/uuid"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{Db: db}
}

func (c *CardRepo) Create(card *models.CreateCard) error {
	_, err := c.Db.Exec("INSERT INTO card(id, number, user_id) VALUES ($1, $2, $3)",
		uuid.NewString(), card.Number, card.UserId)
	return err
}

func (c *CardRepo) GetById(id string) (*models.Card, error) {
	var card models.Card
	err := c.Db.QueryRow("SELECT id, number, user_id FROM card WHERE id = $1", id).
		Scan(&card.Id, &card.Number, &card.UserId)
	if err != nil {
		return nil, err
	}
	return &card, nil
}

func (c *CardRepo) Update(card *models.Card) error {
	_, err := c.Db.Exec("UPDATE card SET number = $1, user_id = $2 WHERE id = $3",
		card.Number, card.UserId, card.Id)
	return err
}

func (c *CardRepo) Delete(id string) error {
	_, err := c.Db.Exec("DELETE FROM card WHERE id = $1", id)
	return err
}
