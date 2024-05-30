package postgres

import (
	"database/sql"
	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson28/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) GetAllCourses() ([]model.Course, error) {
	rows, err := c.DB.Query(`SELECT id, name, field FROM course`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Course
	for rows.Next() {
		var course model.Course
		err = rows.Scan(&course.ID, &course.Name, &course.Field)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (c *CourseRepo) GetByID(id string) (model.Course, error) {
	var course model.Course
	err := c.DB.QueryRow(`SELECT id, name, field FROM course WHERE id = $1`, id).
		Scan(&course.ID, &course.Name, &course.Field)
	if err != nil {
		return model.Course{}, err
	}
	return course, nil
}

func (c *CourseRepo) Create(course *model.Course) error {
	_, err := c.DB.Exec(`INSERT INTO course (id, name, field) VALUES ($1, $2, $3)`,
		course.ID, course.Name, course.Field)
	return err
}

func (c *CourseRepo) Update(course model.Course) error {
	_, err := c.DB.Exec(`UPDATE course SET name=$1, field=$2 WHERE id=$3`,
		course.Name, course.Field, course.ID)
	return err
}

func (c *CourseRepo) Delete(id string) error {
	_, err := c.DB.Exec(`DELETE FROM course WHERE id=$1`, id)
	return err
}
