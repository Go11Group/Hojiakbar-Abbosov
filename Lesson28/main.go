package main

import (
	"fmt"
	"log"

	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson28/model"
	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson28/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	studentRepo := postgres.NewStudentRepo(db)
	courseRepo := postgres.NewCourseRepo(db)

	// Example usage of studentRepo
	students, err := studentRepo.GetAllStudents()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Students:", students)

	student, err := studentRepo.GetByID("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Student by ID:", student)

	newStudent := model.Student{
		ID:       "4",
		Name:     "New Student",
		Age:      20,
		Gender:   "M",
		CourseID: 1,
	}
	if err := studentRepo.Create(newStudent); err != nil {
		log.Fatal(err)
	}

	// Example usage of courseRepo
	courses, err := courseRepo.GetAllCourses()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Courses:", courses)

	course, err := courseRepo.GetByID("1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Course by ID:", course)

	newCourse := model.Course{
		ID:    "3",
		Name:  "New Course",
		Field: "Science",
	}
	if err := courseRepo.Create(newCourse); err != nil {
		log.Fatal(err)
	}
}
