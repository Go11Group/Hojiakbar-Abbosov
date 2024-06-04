package main

import "fmt"

func sortSlice(slice []int) {
	n := len(slice)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if slice[j] < slice[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			slice[i], slice[minIndex] = slice[minIndex], slice[i]
		}
	}
}

func main() {
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	sortSlice(numbers)
	fmt.Println(numbers) 
}
