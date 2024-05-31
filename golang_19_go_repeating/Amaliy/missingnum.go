package main

import "fmt"

func main() {
	
    nums := []int{1, 2, 3, 4, 6, 7, 8}
    sum := 0
    max := nums[0]
    min := nums[0]

    for _, num := range nums {
        sum += num
        if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }

    expectedSum := (min + max) * (len(nums) + 1) / 2
    missingNumber := expectedSum - sum
    fmt.Println("Missing number:", missingNumber) 
}


// package main

// import "fmt"

// func main() {
//     var n int
//     fmt.Print("Sonlar oralig'ida tushirib qoldirilgan sonning massivini topish uchun n ni kiriting: ")
//     fmt.Scan(&n)

//     nums := make([]int, n-1) // massiv uzunligini aniqlash uchun n-1 elementlar
//     sum := 0
//     max := 0
//     min := n

//     fmt.Println("Tushirib qoldirilgan sonni kiriting:")
//     for i := 0; i < n-1; i++ {
//         fmt.Scan(&nums[i])
//         sum += nums[i]
//         if nums[i] > max {
//             max = nums[i]
//         }
//         if nums[i] < min {
//             min = nums[i]
//         }
//     }
// 	fmt.Println(nums)
//     expectedSum := (min + max) * (n + 1) / 2
//     missingNumber := expectedSum - sum
//     fmt.Println("Tushirib qoldirilgan son:", missingNumber)
// }
