package main

import (
 "fmt"
)

func nextPermutation(nums []int) {
 n := len(nums)
 i := n - 2

 for i >= 0 && nums[i] >= nums[i+1] {
  i--
 }

 if i >= 0 {
  j := n - 1
  for j >= 0 && nums[j] <= nums[i] {
   j--
  }
  
  nums[i], nums[j] = nums[j], nums[i]
 }

 
 reverse(nums, i+1)
}

func reverse(nums []int, start int) {
 i, j := start, len(nums)-1
 for i < j {
  nums[i], nums[j] = nums[j], nums[i]
  i++
  j--
 }
}

func main() {
 nums1 := []int{1, 2, 3}
 nextPermutation(nums1)
 fmt.Println(nums1) 

 nums2 := []int{3, 2, 1}
 nextPermutation(nums2)
 fmt.Println(nums2) 

 nums3 := []int{1, 1, 5}
 nextPermutation(nums3)
 fmt.Println(nums3) 
}