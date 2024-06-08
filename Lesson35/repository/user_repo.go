package repository

import (
	"leetcode-api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return user, err
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}

func (r *UserRepository) Update(user models.User) (models.User, error) {
	err := r.DB.Save(&user).Error
	return user, err
}

func (r *UserRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.User{}, id).Error
	return err
}
