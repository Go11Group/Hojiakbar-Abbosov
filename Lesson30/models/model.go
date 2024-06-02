package models

type User struct {
    ID       int
    Username string
    Email    string 
    Password string
}

type Product struct {
    ID            int
    Name          string
    Description   string
    Price         float64
    StockQuantity int
}
