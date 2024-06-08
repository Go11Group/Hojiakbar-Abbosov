package repository

import (
	"leetcode-api/models"
	"gorm.io/gorm"
)

type ProblemRepository struct {
	DB *gorm.DB
}

func NewProblemRepository(db *gorm.DB) *ProblemRepository {
	return &ProblemRepository{DB: db}
}

func (r *ProblemRepository) FindAll() ([]models.Problem, error) {
	var problems []models.Problem
	err := r.DB.Find(&problems).Error
	return problems, err
}

func (r *ProblemRepository) FindByID(id uint) (models.Problem, error) {
	var problem models.Problem
	err := r.DB.First(&problem, id).Error
	return problem, err
}

func (r *ProblemRepository) Create(problem models.Problem) (models.Problem, error) {
	err := r.DB.Create(&problem).Error
	return problem, err
}

func (r *ProblemRepository) Update(problem models.Problem) (models.Problem, error) {
	err := r.DB.Save(&problem).Error
	return problem, err
}

func (r *ProblemRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Problem{}, id).Error
	return err
}