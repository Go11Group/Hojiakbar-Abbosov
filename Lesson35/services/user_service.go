package services

import (
	"leetcode-api/models"
	"leetcode-api/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.FindAll()
}

func (s *UserService) GetUserByID(id uint) (models.User, error) {
	return s.Repo.FindByID(id)
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.Repo.Create(user)
}

func (s *UserService) UpdateUser(user models.User) (models.User, error) {
	return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repo.Delete(id)
}
