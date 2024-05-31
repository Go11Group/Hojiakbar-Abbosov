package main

import "fmt"

func mostFrequent(slice []int) int {
    hashMap := make(map[int]int)
    maxCount := 0
    var n int

    for _, num := range slice {
        hashMap[num]++
        if hashMap[num] > maxCount {
            maxCount = hashMap[num]
            n = num
        }
    }

    return n
}

func main() {
    slice := []int{1, 2, 3, 4, 5, 6, 3, 7}
    fmt.Println("Eng ko'p uchragan son:", mostFrequent(slice))
}

