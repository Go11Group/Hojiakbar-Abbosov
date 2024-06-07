package postgres

import (
    "database/sql"
    "api/model"
)

type UserRepo struct {
    db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
    return &UserRepo{db: db}
}

func (repo *UserRepo) CreateUser(user *model.User) error {
    err := repo.db.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
        user.Username, user.Email, user.Password).Scan(&user.ID)
    return err
}

func (repo *UserRepo) GetUser(id int) (*model.User, error) {
    user := &model.User{}
    err := repo.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = $1", id).
        Scan(&user.ID, &user.Username, &user.Email, &user.Password)
    if err == sql.ErrNoRows {
        return nil, nil
    }
    return user, err
}

func (repo *UserRepo) UpdateUser(user *model.User) error {
    _, err := repo.db.Exec("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4",
        user.Username, user.Email, user.Password, user.ID)
    return err
}

func (repo *UserRepo) DeleteUser(id int) error {
    _, err := repo.db.Exec("DELETE FROM users WHERE id = $1", id)
    return err
}
