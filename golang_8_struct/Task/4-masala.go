package main

import "fmt"

func reverseSlice(ptr *[]int) {
 	slice := *ptr
 	length := len(slice)
 	for i := 0; i < length/2; i++ {
  		slice[i], slice[length-1-i] = slice[length-1-i], slice[i]
 	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Initial slice:", numbers)

	reverseSlice(&numbers)

	fmt.Println("Reversed slice:", numbers)
}