package repositories

import (
    "database/sql"
    "my_module/models"
)

type ProductRepository struct {
    db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
    return &ProductRepository{db}
}

func (r *ProductRepository) Create(product *models.Product) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4)", product.Name, product.Description, product.Price, product.StockQuantity)
    if err != nil {
        return err
    }

    return tx.Commit()
}

func (r *ProductRepository) Get(id int) (*models.Product, error) {
    var product models.Product
    err := r.db.QueryRow("SELECT id, name, description, price, stock_quantity FROM products WHERE id = $1", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity)
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *ProductRepository) Update(product *models.Product) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5", product.Name, product.Description, product.Price, product.StockQuantity, product.ID)
    if err != nil {
        return err
    }

    return tx.Commit()
}

func (r *ProductRepository) Delete(id int) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("DELETE FROM products WHERE id = $1", id)
    if err != nil {
        return err
    }

    return tx.Commit()
}
