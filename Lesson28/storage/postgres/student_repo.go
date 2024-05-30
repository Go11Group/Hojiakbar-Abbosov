package postgres

import (
	"database/sql"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson28/model"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.Student, error) {
	rows, err := u.Db.Query(`SELECT s.id, s.name, s.age, s.gender, s.course_id FROM student s`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Student
	for rows.Next() {
		var student model.Student
		err = rows.Scan(&student.ID, &student.Name, &student.Age, &student.Gender, &student.CourseID)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (u *StudentRepo) GetByID(id string) (model.Student, error) {
	var student model.Student
	err := u.Db.QueryRow(`SELECT s.id, s.name, s.age, s.gender, s.course_id FROM student s WHERE s.id = $1`, id).
		Scan(&student.ID, &student.Name, &student.Age, &student.Gender, &student.CourseID)
	if err != nil {
		return model.Student{}, err
	}
	return student, nil
}

func (u *StudentRepo) Create(student model.Student) error {
	_, err := u.Db.Exec(`INSERT INTO student (id, name, age, gender, course_id) 
		VALUES ($1, $2, $3, $4, $5)`, student.ID, student.Name, student.Age, student.Gender, student.CourseID)
	return err
}

func (u *StudentRepo) Update(student model.Student) error {
	_, err := u.Db.Exec(`UPDATE student SET name=$1, age=$2, gender=$3, course_id=$4 WHERE id=$5`,
		student.Name, student.Age, student.Gender, student.CourseID, student.ID)
	return err
}

func (u *StudentRepo) Delete(id string) error {
	_, err := u.Db.Exec(`DELETE FROM student WHERE id=$1`, id)
	return err
}
