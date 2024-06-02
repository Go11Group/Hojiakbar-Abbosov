package main

import (
    "fmt"
    "log"
    "my_module/db"
    "my_module/models"
    "my_module/repositories"
)

func main() {
    database, err := db.ConnectDB()
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    userRepo := repositories.NewUserRepository(database)
    productRepo := repositories.NewProductRepository(database)
    userProductRepo := repositories.NewUserProductRepository(database)

    err = userRepo.Create(&models.User{Username: "john_doe", Email: "john@example.com", Password: "securepassword"})
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
    }

    err = productRepo.Create(&models.Product{Name: "Laptop", Description: "A powerful laptop", Price: 1200.00, StockQuantity: 10})
    if err != nil {
        log.Fatalf("Failed to create product: %v", err)
    }

    err = userProductRepo.Create(1, 1)
    if err != nil {
        log.Fatalf("Failed to link user to product: %v", err)
    }

    user, err := userRepo.Get(1)
    if err != nil {
        log.Fatalf("Failed to get user: %v", err)
    }
    fmt.Printf("User: %+v\n", user)

    product, err := productRepo.Get(1)
    if err != nil {
        log.Fatalf("Failed to get product: %v", err)
    }
    fmt.Printf("Product: %+v\n", product)

    err = userRepo.Update(&models.User{ID: 1, Username: "john_doe_updated", Email: "john_updated@example.com", Password: "newsecurepassword"})
    if err != nil {
        log.Fatalf("Failed to update user: %v", err)
    }

    err = productRepo.Update(&models.Product{ID: 1, Name: "Laptop Updated", Description: "An even more powerful laptop", Price: 1300.00, StockQuantity: 8})
    if err != nil {
        log.Fatalf("Failed to update product: %v", err)
    }

    err = userRepo.Delete(1)
    if err != nil {
        log.Fatalf("Failed to delete user: %v", err)
    }

    err = productRepo.Delete(1)
    if err != nil {
        log.Fatalf("Failed to delete product: %v", err)
    }
}
