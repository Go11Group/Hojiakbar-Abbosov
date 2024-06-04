package main

import (
	"fmt"
	"sync"
)

func calculateSum(start, end int, wg *sync.WaitGroup, totalSum *int) {
	defer wg.Done()
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	*totalSum += sum
}

func main() {
	var wg sync.WaitGroup
	totalSum := 0

	for i := 0; i < 10; i++ {
		start := i*100 + 1
		end := (i + 1) * 100
		wg.Add(1)
		go calculateSum(start, end, &wg, &totalSum)
	}

	wg.Wait()

	fmt.Println("Sum of numbers from 1 to 1000:", totalSum)
}
