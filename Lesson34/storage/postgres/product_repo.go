package postgres

import (
    "database/sql"
    "api/model"
)

type ProductRepo struct {
    db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
    return &ProductRepo{db: db}
}

func (repo *ProductRepo) CreateProduct(product *model.Product) error {
    err := repo.db.QueryRow("INSERT INTO products (name, description, price, stock_quantity) VALUES ($1, $2, $3, $4) RETURNING id",
        product.Name, product.Description, product.Price, product.StockQuantity).Scan(&product.ID)
    return err
}

func (repo *ProductRepo) GetProduct(id int) (*model.Product, error) {
    product := &model.Product{}
    err := repo.db.QueryRow("SELECT id, name, description, price, stock_quantity FROM products WHERE id = $1", id).
        Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.StockQuantity)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return product, err
}

func (repo *ProductRepo) UpdateProduct(product *model.Product) error {
    _, err := repo.db.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5",
        product.Name, product.Description, product.Price, product.StockQuantity, product.ID)
    return err
}

func (repo *ProductRepo) DeleteProduct(id int) error {
    _, err := repo.db.Exec("DELETE FROM products WHERE id = $1", id)
    return err
}
