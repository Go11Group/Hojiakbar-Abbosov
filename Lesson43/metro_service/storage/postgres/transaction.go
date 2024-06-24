package postgres

import (
	"database/sql"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/metro_service/models"
	"github.com/google/uuid"
)

type TransactionRepo struct {
	Db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{Db: db}
}

func (t *TransactionRepo) Create(transaction *models.CreateTransaction) error {
	_, err := t.Db.Exec("INSERT INTO transaction(id, card_id, amount, terminal_id, transaction_type) VALUES ($1, $2, $3, $4, $5)",
		uuid.NewString(), transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType)
	return err
}

func (t *TransactionRepo) GetById(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := t.Db.QueryRow("SELECT id, card_id, amount, terminal_id, transaction_type FROM transaction WHERE id = $1", id).
		Scan(&transaction.Id, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (t *TransactionRepo) Update(transaction *models.Transaction) error {
	_, err := t.Db.Exec("UPDATE transaction SET card_id = $1, amount = $2, terminal_id = $3, transaction_type = $4 WHERE id = $5",
		transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType, transaction.Id)
	return err
}

func (t *TransactionRepo) Delete(id string) error {
	_, err := t.Db.Exec("DELETE FROM transaction WHERE id = $1", id)
	return err
}
