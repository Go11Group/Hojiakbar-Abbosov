package main

import (
 "fmt"
)

func squareElements(slice []*int) {
 	for _, n := range slice {
  		*n *= *n 
 	}
}

func main() {
 	numbers := []*int{new(int), new(int), new(int)}
 	*numbers[0] = 2
 	*numbers[1] = 3
 	*numbers[2] = 4

 	fmt.Println("Initial slice:")
 	for _, n := range numbers {
  		fmt.Print(*n, " ")
 	}
 	fmt.Println()

 	squareElements(numbers)

 	fmt.Println("Updated slice:")
 	for _, n := range numbers {
  		fmt.Print(*n, " ")
 	}
 	fmt.Println()
}