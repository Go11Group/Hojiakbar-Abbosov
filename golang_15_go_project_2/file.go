package main

import (
 "encoding/json"
 "fmt"
 "os"
)

type Company struct {
	ID      int    
	Name    string 
	Phone   string 
	Address string 
}

func CreateCompany(id int, name, phone, address string) *Company {
	return &Company{
	ID:      id,
	Name:    name,
	Phone:   phone,
	Address: address,
 }
}

func main() {
	var id int
	var name, phone, address string

	fmt.Println("Enter company ID:")
	fmt.Scanln(&id)

	fmt.Println("Enter company name:")
	fmt.Scanln(&name)

	fmt.Println("Enter company phone number:")
	fmt.Scanln(&phone)

	fmt.Println("Enter company address:")
	fmt.Scanln(&address)

	company := CreateCompany(id, name, phone, address)
	fmt.Println("Company:", company)

 	jsonData, err := json.MarshalIndent(company, "", "    ")
 		if err != nil {
  			fmt.Println("Error encoding JSON:", err)
  			return
 }

 	file, err := os.Create("company.json")
 	if err != nil {
  		fmt.Println("Error creating file:", err)
  		return
 }
 	defer file.Close()

 	_, err = file.Write(jsonData)
 	if err != nil {
  		fmt.Println("Error writing JSON to file:", err)
  		return
 }

 	fmt.Println("Company data has been written to company.json")
}