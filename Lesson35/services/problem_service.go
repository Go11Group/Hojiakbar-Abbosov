package services

import (
	"leetcode-api/models"
	"leetcode-api/repository"
)

type ProblemService struct {
	Repo *repository.ProblemRepository
}

func NewProblemService(repo *repository.ProblemRepository) *ProblemService {
	return &ProblemService{Repo: repo}
}

func (s *ProblemService) GetAllProblems() ([]models.Problem, error) {
	return s.Repo.FindAll()
}

func (s *ProblemService) GetProblemByID(id uint) (models.Problem, error) {
	return s.Repo.FindByID(id)
}

func (s *ProblemService) CreateProblem(problem models.Problem) (models.Problem, error) {
	return s.Repo.Create(problem)
}

func (s *ProblemService) UpdateProblem(problem models.Problem) (models.Problem, error) {
	return s.Repo.Update(problem)
}

func (s *ProblemService) DeleteProblem(id uint) error {
	return s.Repo.Delete(id)
}
