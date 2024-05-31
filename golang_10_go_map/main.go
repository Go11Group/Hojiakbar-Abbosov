package main

import "fmt"

func main() {
  slice := []int{1,2,7,4,5,6,3,8,9,3}
  var hashMap = map[int]int{23:4}
  for _,n := range slice {
    hashMap [n]++
    if hashMap [n]  > 1{
      fmt.Println(n)
      break
    }
  }
}
