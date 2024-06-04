package main

import (
	"fmt"
	
)

func main() {
	slice := []int{10, 30, 50, 20, 40, 70, 90, 60, 80, 100}

	k := 1
	count := 0
	for _, num := range slice {
		if num > k {
			count++
		}
	}

	fmt.Printf("Slice ichida %d dan kichik elementlar soni: %d\n", k, count)
}
