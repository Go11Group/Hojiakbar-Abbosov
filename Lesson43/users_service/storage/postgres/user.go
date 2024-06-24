package postgres

import (
	"database/sql"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson43/users_service/models"
	"github.com/google/uuid"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) Create(user *models.CreateUser) error {
	_, err := u.Db.Exec("INSERT INTO users(id, name, phone, age) VALUES ($1, $2, $3, $4)",
		uuid.NewString(), user.Name, user.Phone, user.Age)
	return err
}

func (u *UserRepo) GetById(id string) (*models.User, error) {
	var user  = models.User{Id : id}
	err := u.Db.QueryRow("SELECT id, name, phone, age FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.Name, &user.Phone, &user.Age)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) Update(user *models.UpdateUser) error {
	_, err := u.Db.Exec("UPDATE users SET name = $1, phone = $2, age = $3 WHERE id = $4",
		user.Name, user.Phone, user.Age, user.Id)
	return err
}

func (u *UserRepo) Delete(id string) error {
	_, err := u.Db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
