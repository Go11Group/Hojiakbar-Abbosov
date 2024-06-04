package main

import (
	"fmt"
)

func main() {

	slice1 := []int{3, 7, 4, 3, 9}
	slice2 := []int{5, 1, 8, 2, 6}

	fmt.Println("1-slice:", slice1)
	fmt.Println("2-slice:", slice2)

	found := false
	for _, num1 := range slice1 {
		for _, num2 := range slice2 {
			if num1 == num2 {
				fmt.Println("Bir xil son:", num1)
				found = true
				
			}
		}
		
	}

	if !found {
		fmt.Println("Uchraydigan bir xil son topilmadi")
	}
}
