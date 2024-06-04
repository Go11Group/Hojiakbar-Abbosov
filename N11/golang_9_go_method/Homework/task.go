package main

import "fmt"

type Employee struct {
	ID           int
	Name         string
	DepartmentID int
	Projects     []string
}

type Department struct {
	ID        int
	Name      string
	Employees []Employee
}

type Project struct {
	ID          int
	Name        string
	DepartmentID int
	Employees   []Employee
}

func (e *Employee) AssignToProject(projectID int) {
	e.Projects = append(e.Projects, fmt.Sprintf("Project-%d", projectID))
}

func (d *Department) AddEmployee(employee Employee) {
	d.Employees = append(d.Employees, employee)
}

func (d *Department) GetEmployees() []Employee {
	return d.Employees
}

func (p *Project) AddEmployee(employee Employee) {
	p.Employees = append(p.Employees, employee)
}

func main() {
	emp1 := Employee{ID: 1, Name: "John Doe", DepartmentID: 101}
	emp2 := Employee{ID: 2, Name: "Jane Smith", DepartmentID: 102}
	emp3 := Employee{ID: 3, Name: "Tom Brown", DepartmentID: 101}

	dept1 := Department{ID: 101, Name: "Engineering"}
	dept2 := Department{ID: 102, Name: "Marketing"}

	dept1.AddEmployee(emp1)
	dept1.AddEmployee(emp3)
	dept2.AddEmployee(emp2)

	project1 := Project{ID: 1, Name: "Project A", DepartmentID: 101}
	project2 := Project{ID: 2, Name: "Project B", DepartmentID: 102}

	emp1.AssignToProject(project1.ID)
	emp2.AssignToProject(project2.ID)
	emp3.AssignToProject(project1.ID)

	fmt.Println("Employees in Engineering Department:")
	engineeringEmployees := dept1.GetEmployees()
	for _, emp := range engineeringEmployees {
		fmt.Println(emp.Name)
	}

	fmt.Println("\nEmployees in Project A:")
	projectAEmployees := project1.Employees
	fmt.Println(projectAEmployees)
	for _, emp := range projectAEmployees {
		fmt.Println(emp.Name)
	}
}
