package main

import (
    "language-learning-app/database"
    "language-learning-app/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    database.Connect()

    routes.SetupRoutes(r)
    r.Run(":8080")
}
