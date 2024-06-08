package repository

import (
	"leetcode-api/models"
	"gorm.io/gorm"
)

type SolvedProblemRepository struct {
	DB *gorm.DB
}

func NewSolvedProblemRepository(db *gorm.DB) *SolvedProblemRepository {
	return &SolvedProblemRepository{DB: db}
}

func (r *SolvedProblemRepository) FindAll() ([]models.SolvedProblem, error) {
	var solvedProblems []models.SolvedProblem
	err := r.DB.Find(&solvedProblems).Error
	return solvedProblems, err
}

func (r *SolvedProblemRepository) FindByID(id uint) (models.SolvedProblem, error) {
	var solvedProblem models.SolvedProblem
	err := r.DB.First(&solvedProblem, id).Error
	return solvedProblem, err
}

func (r *SolvedProblemRepository) Create(solvedProblem models.SolvedProblem) (models.SolvedProblem, error) {
	err := r.DB.Create(&solvedProblem).Error
	return solvedProblem, err
}

func (r *SolvedProblemRepository) Update(solvedProblem models.SolvedProblem) (models.SolvedProblem, error) {
	err := r.DB.Save(&solvedProblem).Error
	return solvedProblem, err
}

func (r *SolvedProblemRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.SolvedProblem{}, id).Error
	return err
}
