package main

import (
	"fmt"
)

func numIdenticalPairs(nums []int) int {
	hashMap := make(map[int]int)
	
	n := 0
	
	for _, num := range nums {
		hashMap[num]++
	}
	
	for _, i := range hashMap {
		n += i * (i - 1) / 2
	}
	
	return n
}

func main() {

	nums1 := []int{1, 2, 3, 1, 1, 3}
	fmt.Println("Example 1 Output:", numIdenticalPairs(nums1)) 

	nums2 := []int{1, 1, 1, 1}
	fmt.Println("Example 2 Output:", numIdenticalPairs(nums2)) 

	nums3 := []int{1, 2, 3}
	fmt.Println("Example 3 Output:", numIdenticalPairs(nums3)) 
}
