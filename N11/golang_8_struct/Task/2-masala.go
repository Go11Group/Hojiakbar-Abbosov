package main

import "fmt"

func main() {
	a := 10
	n := 90
	b := &a
	c := &b
	d := &c

	fmt.Println("Initial values:")
	fmt.Println("a:", a)
	fmt.Println("b:", *b) 
	fmt.Println("c:", **c) 
	fmt.Println("d:", ***d) 

	**c = n


	fmt.Println("\nUpdated values:")
	fmt.Println("a:", a) 
	fmt.Println("b:", *b) 
	fmt.Println("c:", **c) 
	fmt.Println("d:", ***d) 
}