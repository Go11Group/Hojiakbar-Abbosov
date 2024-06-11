package controllers

import (
	"database/sql"
	"log"
)

type Problem struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetProblems(db *sql.DB) ([]Problem, error) {
	rows, err := db.Query("SELECT id, title, content FROM problems")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []Problem
	for rows.Next() {
		var problem Problem
		if err := rows.Scan(&problem.ID, &problem.Title, &problem.Content); err != nil {
			log.Println(err)
			continue
		}
		problems = append(problems, problem)
	}
	return problems, nil
}

func GetProblem(db *sql.DB, id int) (*Problem, error) {
	var problem Problem
	err := db.QueryRow("SELECT id, title, content FROM problems WHERE id=$1", id).Scan(&problem.ID, &problem.Title, &problem.Content)
	if err != nil {
		return nil, err
	}
	return &problem, nil
}

func CreateProblem(db *sql.DB, problem *Problem) error {
	err := db.QueryRow("INSERT INTO problems (title, content) VALUES ($1, $2) RETURNING id", problem.Title, problem.Content).Scan(&problem.ID)
	return err
}

func UpdateProblem(db *sql.DB, problem *Problem) error {
	_, err := db.Exec("UPDATE problems SET title=$1, content=$2 WHERE id=$3", problem.Title, problem.Content, problem.ID)
	return err
}

func DeleteProblem(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM problems WHERE id=$1", id)
	return err
}
