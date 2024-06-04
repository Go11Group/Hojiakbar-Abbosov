package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	
	connStr := "user=postgres dbname=go11 sslmode=disable password=root"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	defer db.Close()

	file, err := os.Open("query.sql")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error stating file: %v\n", err)
		return
	}

	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	queryList := [3]string{
		`CREATE TABLE IF NOT EXISTS product (
			id UUID PRIMARY KEY,
			name VARCHAR,
			category VARCHAR,
			cost INT
		);`,
		"CREATE INDEX idx_product_name ON product (name);",
		"CREATE INDEX idx_product_category ON product (category);",
	}

	for _, query := range queryList {
		_, err := db.Exec(query)
		if err != nil {
			fmt.Printf("Error executing query: %v\nQuery: %s\n", err, query)
			return
		}
	}

	numRecords := 1000000
	start := time.Now()
	for i := 0; i < numRecords; i++ {
		id := uuid.New()
		name := faker.Name()
		category := faker.Word()
		cost := rand.Intn(10000)
		_, err := db.Exec("INSERT INTO product (id, name, category, cost) VALUES ($1, $2, $3, $4)", id, name, category, cost)
		if err != nil {
			fmt.Printf("Error inserting record: %v\n", err)
			return
		}
	}
	insertionTime := time.Since(start)

	start = time.Now()
	rows, err := db.Query("SELECT * FROM product WHERE name = $1", "John Doe")
	if err != nil {
		fmt.Printf("Error querying database: %v\n", err)
		return
	}
	for rows.Next() {
		var id uuid.UUID
		var name, category string
		var cost int
		if err := rows.Scan(&id, &name, &category, &cost); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return
		}
	}
	nameSearchTime := time.Since(start)

	start = time.Now()
	rows, err = db.Query("SELECT * FROM product WHERE category = $1", "electronics")
	if err != nil {
		fmt.Printf("Error querying database: %v\n", err)
		return
	}
	for rows.Next() {
		var id uuid.UUID
		var name, category string
		var cost int
		if err := rows.Scan(&id, &name, &category, &cost); err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			return
		}
	}
	categorySearchTime := time.Since(start)

	fmt.Printf("Insertion Time: %v\n", insertionTime)
	fmt.Printf("Name Search Time: %v\n", nameSearchTime)
	fmt.Printf("Category Search Time: %v\n", categorySearchTime)
}
