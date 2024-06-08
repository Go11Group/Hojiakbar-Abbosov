package services

import (
	"leetcode-api/models"
	"leetcode-api/repository"
)

type SolvedProblemService struct {
	Repo *repository.SolvedProblemRepository
}

func NewSolvedProblemService(repo *repository.SolvedProblemRepository) *SolvedProblemService {
	return &SolvedProblemService{Repo: repo}
}

func (s *SolvedProblemService) GetAllSolvedProblems() ([]models.SolvedProblem, error) {
	return s.Repo.FindAll()
}

func (s *SolvedProblemService) GetSolvedProblemByID(id uint) (models.SolvedProblem, error) {
	return s.Repo.FindByID(id)
}

func (s *SolvedProblemService) CreateSolvedProblem(solvedProblem models.SolvedProblem) (models.SolvedProblem, error) {
	return s.Repo.Create(solvedProblem)
}

func (s *SolvedProblemService) UpdateSolvedProblem(solvedProblem models.SolvedProblem) (models.SolvedProblem, error) {
	return s.Repo.Update(solvedProblem)
}

func (s *SolvedProblemService) DeleteSolvedProblem(id uint) error {
	return s.Repo.Delete(id)
}
