package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Quote struct {
	ID     int    `json:"id"`
	Text   string `json:"quote"`
	Author string `json:"author"`
}

func main() {
	filePath := "data.json"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	var quotes []Quote
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&quotes)
	if err != nil {
		fmt.Println("Error decoding JSON data:", err)
		return
	}

	for _, q := range quotes {
		fmt.Printf("ID: %d\nQuote: %s\nAuthor: %s\n\n", q.ID, q.Text, q.Author)
	}
}
