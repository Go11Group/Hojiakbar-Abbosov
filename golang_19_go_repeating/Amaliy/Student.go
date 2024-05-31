package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Student struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Courses []Course  `json:"courses"`
	Address Address   `json:"address"`
}

type Course struct {
	CourseID  string `json:"courseId"`
	Grade     string `json:"grade"`
	Professor string `json:"professor"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func writeStudentsToFile(students []Student, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(students); err != nil {
		return err
	}

	return nil
}

func readStudentsFromFile(filename string) ([]Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var students []Student
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&students); err != nil {
		return nil, err
	}

	return students, nil
}

func printStudentDetails(students []Student) {
	for _, student := range students {
		fmt.Printf("ID: %s\nName: %s\nAddress: %s, %s, %s\nCourses:\n", student.ID, student.Name, student.Address.Street, student.Address.City, student.Address.Country)
		for _, course := range student.Courses {
			fmt.Printf("  Course ID: %s\n  Grade: %s\n  Professor: %s\n", course.CourseID, course.Grade, course.Professor)
		}
		fmt.Println()
	}
}

func findTopScoringStudent(students []Student) (*Student, error) {
	if len(students) == 0 {
		return nil, errors.New("no students found")
	}

	var topStudent *Student
	topGrade := -1

	for _, student := range students {
		for _, course := range student.Courses {
			if course.Grade == "A" && topGrade < 5 { // Assuming 'A' is the top grade
				topStudent = &student
				topGrade = 5
			}
		}
	}

	if topStudent == nil {
		return nil, errors.New("no top scoring student found")
	}

	return topStudent, nil
}

func groupStudentsByCategory(students []Student, category string) (map[string][]Student, error) {
	grouped := make(map[string][]Student)

	for _, student := range students {
		switch category {
		case "courses":
			for _, course := range student.Courses {
				if _, ok := grouped[course.CourseID]; !ok {
					grouped[course.CourseID] = []Student{student}
				} else {
					grouped[course.CourseID] = append(grouped[course.CourseID], student)
				}
			}
		case "address":
			addressKey := fmt.Sprintf("%s, %s, %s", student.Address.Street, student.Address.City, student.Address.Country)
			if _, ok := grouped[addressKey]; !ok {
				grouped[addressKey] = []Student{student}
			} else {
				grouped[addressKey] = append(grouped[addressKey], student)
			}
		default:
			return nil, errors.New("invalid category")
		}
	}

	return grouped, nil
}

func main() {
	
	student1 := Student{
		ID:   "123",
		Name: "John Doe",
		Courses: []Course{
			{CourseID: "101", Grade: "A", Professor: "Dr. Smith"},
			{CourseID: "102", Grade: "B", Professor: "Dr. Johnson"},
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			Country: "USA",
		},
	}
	student2 := Student{
		ID:   "456",
		Name: "Jane Doe",
		Courses: []Course{
			{CourseID: "101", Grade: "B", Professor: "Dr. Smith"},
			{CourseID: "103", Grade: "A", Professor: "Dr. Johnson"},
		},
		Address: Address{
			Street:  "456 Elm St",
			City:    "Othertown",
			Country: "USA",
		},
	}

	students := []Student{student1, student2}

	err := writeStudentsToFile(students, "students.json")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Students written to file successfully.")

	readStudents, err := readStudentsFromFile("students.json")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("Students read from file:")
	printStudentDetails(readStudents)

	topStudent, err := findTopScoringStudent(readStudents)
	if err != nil {
		fmt.Println("Error finding top scoring student:", err)
	} else {
		fmt.Println("Top Scoring Student:")
		fmt.Printf("ID: %s\nName: %s\n", topStudent.ID, topStudent.Name)
	}

	courseGroups, err := groupStudentsByCategory(readStudents, "courses")
	if err != nil {
		fmt.Println("Error grouping students by courses:", err)
		return
	}
	fmt.Println("Students grouped by courses:")
	for courseID, students := range courseGroups {
		fmt.Println("Course ID:", courseID)
		fmt.Println("Students:")
		for _, student := range students {
			fmt.Printf("  ID: %s, Name: %s\n", student.ID, student.Name)
		}
	}

	addressGroups, err := groupStudentsByCategory(readStudents, "address")
	if err != nil {
		fmt.Println("Error grouping students by address:", err)
		return
	}
	fmt.Println("Students grouped by address:")
	for addressKey, students := range addressGroups {
		fmt.Println("Address:", addressKey)
		fmt.Println("Students:")
		for _, student := range students {
			fmt.Printf("  ID: %s, Name: %s\n", student.ID, student.Name)
		}
	}
}
