package models

type SolvedProblem struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	UserID    uint `json:"user_id"`
	ProblemID uint `json:"problem_id"`
}