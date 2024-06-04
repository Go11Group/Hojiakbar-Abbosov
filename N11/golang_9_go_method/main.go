package main

import (
 "fmt"
)

type Person struct {
	Name    string
	Age     int
	Country string
}

func fillStructs(slice []Person) {
 	for i := range slice {
		var name string
		var age int
		var country string

		fmt.Printf("Enter name for person %d: ", i+1)
		fmt.Scanln(&name)
		fmt.Printf("Enter age for person %d: ", i+1)
		fmt.Scanln(&age)
		fmt.Printf("Enter country for person %d: ", i+1)
		fmt.Scanln(&country)

		slice[i] = Person{Name: name, Age: age, Country: country}
 }
}

func main() {
	
	var n int
	fmt.Print("Enter the number of persons: ")
	fmt.Scanln(&n)

	persons := make([]Person, n)

	fillStructs(persons)

	fmt.Println("\nFilled structs:")
	for i, p := range persons {
		fmt.Printf("Person %d: Name = %s, Age = %d, Country = %s\n", i+1, p.Name, p.Age, p.Country)
 }
}