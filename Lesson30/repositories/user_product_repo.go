package repositories

import "database/sql"

type UserProductRepository struct {
    db *sql.DB
}

func NewUserProductRepository(db *sql.DB) *UserProductRepository {
    return &UserProductRepository{db}
}

func (r *UserProductRepository) Create(userID, productID int) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("INSERT INTO user_products (user_id, product_id) VALUES ($1, $2)", userID, productID)
    if err != nil {
        return err
    }

    return tx.Commit()
}
