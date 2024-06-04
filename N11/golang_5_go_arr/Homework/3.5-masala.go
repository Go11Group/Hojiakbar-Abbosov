package main

import (
	"fmt"
)
func main() {
	var matn string

	matn  = "Hello world saloM dunyo"

	n := len(matn)

	fmt.Print("Matnning teskari tartibdagi ko'rinishi: ")
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%c", matn[i])
	}
	fmt.Println() 


}
	



