package main

import (
    "log"
    "net/http"
    "os"

    "api/handler"
    "api/storage/postgres"
    "github.com/go-chi/chi/v5"
)

func main() {
    connString := os.Getenv("DATABASE_URL")
    if connString == "" {
        log.Fatal("DATABASE_URL environment variable is not set")
    }

    err := postgres.InitDB(connString)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer postgres.CloseDB()

    userRepo := postgres.NewUserRepo(postgres.GetDB())
    userHandler := &handler.UserHandler{Repo: userRepo}

    productRepo := postgres.NewProductRepo(postgres.GetDB())
    productHandler := &handler.ProductHandler{Repo: productRepo}

    r := chi.NewRouter()

    r.Post("/user", userHandler.CreateUser)
    r.Get("/user/{id}", userHandler.GetUser)
    r.Put("/user/{id}", userHandler.UpdateUser)
    r.Delete("/user/{id}", userHandler.DeleteUser)

    r.Post("/product", productHandler.CreateProduct)
    r.Get("/product/{id}", productHandler.GetProduct)
    r.Put("/product/{id}", productHandler.UpdateProduct)
    r.Delete("/product/{id}", productHandler.DeleteProduct)

    log.Println("Server is running on port 8080")
    http.ListenAndServe(":8080", r)
}
