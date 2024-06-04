package main

import "fmt"

func findMax(slice []int) int {
 	if len(slice) == 0 {
  		return 0 
 }

 	max := slice[0] 
 	for _, value := range slice {
  		if value > max {
   		max = value 
  		}
	}
 	return max
}

func main() {
	slice := []int{8, 11, 15, 9, 1, 25}
	maxNum := findMax(slice)
	fmt.Println("Eng katta son:", maxNum) 
}