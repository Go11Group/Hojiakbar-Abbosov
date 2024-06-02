package repositories

import (
    "database/sql"
    "my_module/models"
)

type UserRepository struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
    return &UserRepository{db}
}

func (r *UserRepository) Create(user *models.User) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
    if err != nil {
        return err
    }

    return tx.Commit()
}

func (r *UserRepository) Get(id int) (*models.User, error) {
    var user models.User
    err := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4", user.Username, user.Email, user.Password, user.ID)
    if err != nil {
        return err
    }

    return tx.Commit()
}

func (r *UserRepository) Delete(id int) error {
    tx, err := r.db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()

    _, err = tx.Exec("DELETE FROM users WHERE id = $1", id)
    if err != nil {
        return err
    }

    return tx.Commit()
}
