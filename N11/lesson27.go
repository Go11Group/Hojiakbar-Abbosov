package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "employee"
	password = "root"
)


func main() {

	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    query := `
        SELECT e.employee_id, e.name, d.department_name, e.age
        FROM Employees e
        JOIN Departments d ON e.department_id = d.department_id
    `
    rows, err := db.Query(query)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    fmt.Println("All employees and their departments:")
    for rows.Next() {
        var employeeID int
        var name, departmentName string
        var age int
        err := rows.Scan(&employeeID, &name, &departmentName, &age)
        if err != nil {
            panic(err)
        }
        fmt.Printf("employee_id: %d, name: %s, department: %s, age: %d\n", employeeID, name, departmentName, age)
    }

    err = rows.Err()
    if err != nil {
        panic(err)
    }

    var employeeID int = 2 
    queryRow := `
        SELECT e.employee_id, e.name, d.department_name, e.age
        FROM Employees e
        JOIN Departments d ON e.department_id = d.department_id
        WHERE e.employee_id = $1
    `
    var name, departmentName string
    var age int
    err = db.QueryRow(queryRow, employeeID).Scan(&employeeID, &name, &departmentName, &age)
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No rows were returned!")
        } else {
            panic(err)
        }
    } else {
        fmt.Printf("Details of employee_id %d: name: %s, department: %s, age: %d\n", employeeID, name, departmentName, age)
    }
}


// db.Query funksiyasi Go dasturlash tilida database/sql paketining bir qismi bo'lib, u SQL so'rovini bajarish 
// va natija satrlarini olish uchun ishlatiladi. Bu funksiya ko'p satrli natijalarni qaytaradi 
// va rows ob'ektini hosil qiladi, bu orqali natija satrlari bo'ylab iteratsiya qilish mumkin.

//SQL so'rovini yozish:
//SQL so'rovini query string o'zgaruvchisida yozib olinadi.

//So'rovni bajarish:
//db.Query funksiyasi chaqiriladi va natijalar rows ob'ektiga saqlanadi.

//Natijalar bo'ylab iteratsiya qilish:
//rows.Next() metodi yordamida har bir natija satriga o'tiladi va rows.Scan yordamida bu satrning ustun qiymatlari olinadi.

//Yopish va xatolarni tekshirish:
//Iteratsiyadan keyin rows.Close() yordamida resurslar tozalanadi va xatolar rows.Err() yordamida tekshiriladi./