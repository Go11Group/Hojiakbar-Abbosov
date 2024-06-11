package main

import (
	"go-gin-crud/database"
	"go-gin-crud/routes"
)

func main() {
	
	database.Connect()
	defer database.Close()

	r := routes.SetupRouter()
	r.Run(":8080")
}
