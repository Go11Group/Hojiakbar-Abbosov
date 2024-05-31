package main

import (
 "fmt"
)

func updatePointerValue(ptr *int) {
 	var n int
 	fmt.Print("Enter a new value: ")
 	fmt.Scanln(&n)
 	*ptr = n
}

func main() {
 	var n int
 	ptr := &n

 	fmt.Println("Initial value of pointer:", *ptr)

 	updatePointerValue(ptr)

 	fmt.Println("Updated value of pointer:", *ptr)
}