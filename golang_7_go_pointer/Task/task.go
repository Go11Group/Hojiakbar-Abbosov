// package main

// import "fmt"

// 1 - masala

// func squares(slice []int) []int {
//  	squared := make([]int, len(slice))
//  	for i, v := range slice {
//   		squared[i] = v * v
//  }
//  	return squared
// }

// func main() {
// 	input := []int{-1, -2, -9, 4, 3, 0}
//  	output := squares(input)
//  	fmt.Println(output) 
// }

// 2 - masala

// func findSmallestPositive(slice []int) int {
//     minPos := -1 
//  	for _, v := range slice {
//   		if v > 0 && (minPos == -1 || v < minPos) {
//    			minPos = v
//   		}
//  	}
//  		return minPos
// }

// func main() {
//  	input := []int{-1, -2, -9, 4, 3, 1, 2, 10}
//  	result := findSmallestPositive(input)
//  	if result == -1 {
//   		fmt.Println(0) 
//  	} else {
//   		fmt.Println(result) 
//  	}
// }

// 3 - masala


// func average(slice []float64) float64 {
//  	var sum float64 = 0
//  	for _, value := range slice {
//   		sum += value
//  	}
//  	return sum / float64(len(slice))
// }

// func main() {
//  	input := []float64{1, 4, 8, 9, 1, 1, 2, 3, 4}
//  	fmt.Printf("O'rta arifmetik qiymat: %.1f\n", average(input)) 
// }

// 4 - masala

// func countMinElements(slice []int) int {
//  	if len(slice) == 0 {
//   		return 0 
//  	}

//  	min := slice[0]
//  	count := 1

//  	for _, value := range slice[1:] {
//   		if value < min {
//    			min = value 
//    			count = 1    
//   		} else if value == min {
//    			count++ 
//   		}
//  	}

//  	return count 
// }

// func main() {
//  	input := []int{1, 4, 8, 9, 1, 1, 2, 3, 4}
//  	fmt.Println("Minimal elementning soni:", countMinElements(input)) 
// }

// 5 - masala

// func isPrime(number int) bool {
//  	if number <= 1 {
//   		return false
//  	}
//  	for i := 2; i*i <= number; i++ {
//   		if number%i == 0 {
//    			return false
//   		}
//  	}
//  	return true
// }

// func printPrimes(N int) {
//  	for number := 2; number <= N; number++ {
//   		if isPrime(number) {
//    			fmt.Printf("%d ", number)
//   		}
//  	}
//  	fmt.Println()
// }

// func main() {
//  	var N int
//  	fmt.Print("N natural sonini kiriting: ")
//  	fmt.Scan(&N)

//  	fmt.Printf("%d gacha bo'lgan tub sonlar: ", N)
//  	printPrimes(N)
// }

// 6 - masala

// func sumDigits(number int) int {
//  	sum := 0
//  	for number > 0 {
//   		sum += number % 10
//   		number /= 10
// 	}
//  	return sum
// }

// func reduceToOneDigit(number int) int {
//  	for number >= 10 {
//   		number = sumDigits(number)
//  	}
//  	return number
// }

// func main() {
//  	var number int
//  	fmt.Print("Sonni kiriting: ")
//  	fmt.Scan(&number)

//  	result := reduceToOneDigit(number)
//  	fmt.Printf("Natija: %d\n", result)
// }

