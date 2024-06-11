package models

import (
	"database/sql"
	"log"
)

type SolvedProblem struct {
	ID        int `json:"id"`
	UserID    int `json:"user_id"`
	ProblemID int `json:"problem_id"`
}

func GetSolvedProblems(db *sql.DB) ([]SolvedProblem, error) {
	rows, err := db.Query("SELECT id, user_id, problem_id FROM solved_problems")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var solvedProblems []SolvedProblem
	for rows.Next() {
		var solvedProblem SolvedProblem
		if err := rows.Scan(&solvedProblem.ID, &solvedProblem.UserID, &solvedProblem.ProblemID); err != nil {
			log.Println(err)
			continue
		}
		solvedProblems = append(solvedProblems, solvedProblem)
	}
	return solvedProblems, nil
}

func GetSolvedProblem(db *sql.DB, id int) (*SolvedProblem, error) {
	var solvedProblem SolvedProblem
	err := db.QueryRow("SELECT id, user_id, problem_id FROM solved_problems WHERE id=$1", id).Scan(&solvedProblem.ID, &solvedProblem.UserID, &solvedProblem.ProblemID)
	if err != nil {
		return nil, err
	}
	return &solvedProblem, nil
}

func CreateSolvedProblem(db *sql.DB, solvedProblem *SolvedProblem) error {
	err := db.QueryRow("INSERT INTO solved_problems (user_id, problem_id) VALUES ($1, $2) RETURNING id", solvedProblem.UserID, solvedProblem.ProblemID).Scan(&solvedProblem.ID)
	return err
}

func UpdateSolvedProblem(db *sql.DB, solvedProblem *SolvedProblem) error {
	_, err := db.Exec("UPDATE solved_problems SET user_id=$1, problem_id=$2 WHERE id=$3", solvedProblem.UserID, solvedProblem.ProblemID, solvedProblem.ID)
	return err
}

func DeleteSolvedProblem(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM solved_problems WHERE id=$1", id)
	return err
}
