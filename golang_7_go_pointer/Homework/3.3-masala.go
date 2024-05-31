package main

import "fmt"

func doubleAndAppend(slice []int) []int {
	return append(slice, slice...) 
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = doubleAndAppend(numbers)
	fmt.Println(numbers) 
}
