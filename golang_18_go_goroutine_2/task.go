// 1 - masala

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func user(n int,wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(n)
// }



// func main(){

// 	var wg sync.WaitGroup
// 	var num int 
// 	fmt.Println("N = ")
// 	fmt.Scan(&num)

// 	for i:= 0;i<num;i++{
// 		wg.Add(1)
// 		go user(i,&wg)
// 	}
// 	wg.Wait()
	
// }

// 2- masala

// package main

// import (
//   "fmt"
//   "sync"
// )

// type Rectangle struct {
//  	a, b float64
// }

// func calculateArea(rect Rectangle, wg *sync.WaitGroup) {
//  	defer wg.Done()

//  	area := rect.a * rect.b
//  	fmt.Println("Yuza:", area)
// }

// func calculatePerimeter(rect Rectangle, wg *sync.WaitGroup) {
//  	defer wg.Done()

//  	perimeter := 2 * (rect.a + rect.b)
//  	fmt.Println("Perimetri:", perimeter)
// }

// func main() {

// 	var wg sync.WaitGroup

// 	rect := Rectangle{a: 10, b: 20}

// 	wg.Add(2)
// 	go calculateArea(rect, &wg)
// 	go calculatePerimeter(rect, &wg)

// 	wg.Wait()
// }

// 3 - masala

// package main

// import (
//  "fmt"
//  "sync"
// )

// func add(num1, num2 int, wg *sync.WaitGroup) {
//  	defer wg.Done()

//  	fmt.Println("num1+num2 =", num1+num2)
// }

// func subtract(num1, num2 int, wg *sync.WaitGroup) {
//  	defer wg.Done()

//  	fmt.Println("num1-num2 =", num1-num2)
// }

// func multiply(num1, num2 int, wg *sync.WaitGroup) {
//  	defer wg.Done()

//  	fmt.Println("num1*num2 =", num1*num2)
// }

// func divide(num1, num2 int, wg *sync.WaitGroup) {
//  	defer wg.Done()

//  	fmt.Println("num1/num2 =", num1/num2)
// }

// func main() {

// 	var wg sync.WaitGroup

// 	num1, num2 := 10, 2

// 	wg.Add(4)
// 	go add(num1, num2, &wg)
// 	go subtract(num1, num2, &wg)
// 	go multiply(num1, num2, &wg)
// 	go divide(num1, num2, &wg)

// 	wg.Wait()
// }

// 4 - masala 

// package main

// import (
//  "fmt"
//  "sync"
// )

// func sumNumbers(n int, sum *int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := 1; i <= n; i++ {
// 		*sum += i
//  }
// }

// func main() {

// 	var wg sync.WaitGroup
// 	var sum int


// 	var n int 
// 	fmt.Println("N = ")
// 	fmt.Scan(&n)
// 	wg.Add(1)
// 	go sumNumbers(n, &sum, &wg)

// 	wg.Wait()
// 	fmt.Println("Sum:", sum)
// }

// 5 - masala

// package main

// import (
//  "fmt"
//  "sync"
// )

// func calculateSum(slice []int, result *int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	sum := 0
// 	for _, num := range slice {
// 		sum += num
// 	}
// 		*result = sum
// }

// func main() {
// 	var wg sync.WaitGroup
// 	slc := [][]int{{1, 2, 3}, {4, 5, 6}, {6, 67, 67}}
// 	results := make([]int, len(slc))

//  	for i, subSlice := range slc {
//   		wg.Add(1)
//   		go calculateSum(subSlice, &results[i], &wg)
//  	}
//  	wg.Wait()

//  	fmt.Println("Results:", results)
// }

// 6 - masala

// package main

// import (
//  "fmt"
//  "sort"
// )

// func main() {
//  	slc := [][]int{{1, 2, -89}, {19, 5, 63}, {6, -607, 67}}

//  	for _, subSlice := range slc {
//   		sort.Ints(subSlice)
//   		fmt.Println(subSlice)
//  }
// }


package main

import (
 "fmt"
 "sync"
)

func calculateSumASCII(s string, result *int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, char := range s {
		sum += int(char)
	}
	*result = sum
}

func main() {
	
	var wg sync.WaitGroup
	s := "Salom"

 	if len(s) < 3 {
  		fmt.Println("Xatolik: String 3 ta harfli emas.")
  		return
 	}

 	var results []int
 	for i := 0; i < len(s)-2; i += 3 {
  		wg.Add(1)
  		go calculateSumASCII(s[i:i+3], &results[i/3], &wg)
 	}

 	wg.Wait()
 	totalSum := 0
 	for _, sum := range results {
  		totalSum += sum
 	}
 	fmt.Println("Jami yig'indisi:", totalSum)
}