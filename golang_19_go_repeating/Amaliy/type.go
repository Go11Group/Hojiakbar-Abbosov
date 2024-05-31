package main

import (
	"fmt"
	"strconv"
)

func checkDataType(input string) string {
	if _, err := strconv.Atoi(input); err == nil {
		return "int"
	} else if _, err := strconv.ParseFloat(input, 64); err == nil {
		return "float64"
	}
	return "string"
}

func main() {
	
	var userInput string
	fmt.Print("Qiymat kiriting: ")
	fmt.Scanln(&userInput)

	dataType := checkDataType(userInput)
	fmt.Printf("Kiritilgan ma'lumot %s tipida\n", dataType)
}
