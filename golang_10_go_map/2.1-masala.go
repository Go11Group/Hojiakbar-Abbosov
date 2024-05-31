package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		n := target - num
		if val, ok := hashMap[n]; ok {
			return []int{val, i}
		}
		hashMap[num] = i
	}

	return []int{}
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Println("Example 1 Output:", twoSum(nums1, target1)) 

	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Println("Example 2 Output:", twoSum(nums2, target2)) 

	nums3 := []int{3, 3}
	target3 := 6
	fmt.Println("Example 3 Output:", twoSum(nums3, target3)) 
}
