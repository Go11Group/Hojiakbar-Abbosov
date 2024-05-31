package main

import (
	"fmt"
)

func main() {
	fmt.Print("Matnni kiriting: ")
	var matn string
	fmt.Scanln(&matn)

	n := len(matn)

	fmt.Print("Matnning teskari tartibdagi ko'rinishi: ")
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%c", matn[i])
	}
	fmt.Println() 
}
