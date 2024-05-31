package main

import (
	"fmt"
	"log"
	"module/db"
	"module/models"
	"module/user"
)

func main() {
	
	con := "postgres://postgres:root@localhost:5432/go11?sslmode=disable"

	database, err := db.InitDB(con, &models.User{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = user.CreateUser(database, "John", "Doe", "john.doe@example.com", "password123", 30, "Engineering", "Male", true)
	if err != nil {
		fmt.Println("Error creating user:", err)
	} else {
		fmt.Println("User created successfully")
	}

	users, err := user.ReadUsers(database)
	if err != nil {
		fmt.Println("Error reading users:", err)
	} else {
		fmt.Println("Users:", users)
	}

	err = user.UpdateUser(database, 1, "john.doe@example.com", "newpassword456")
	if err != nil {
		fmt.Println("Error updating user:", err)
	} else {
		fmt.Println("User updated successfully")
	}

	err = user.DeleteUser(database, 1)
	if err != nil {
		fmt.Println("Error deleting user:", err)
	} else {
		fmt.Println("User deleted successfully")
	}
}
