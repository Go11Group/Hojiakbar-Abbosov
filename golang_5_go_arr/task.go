package main

import (
	"fmt"
	"math/rand"
)

func main(){
	// var a,b float64

	// fmt.Scan(&a,&b)

	// for i := a ; i <= b ; i +=1.0{
	// 	for j := 2.0 ; j < i ; j += 1.0{
	// 		if i / j == j{
	// 			fmt.Println(i)
	// 		}
	// 	}
		
	// }
	
	// var nums = [10]int{1,3,9,4,5,6,7,2,10,8}

	// fmt.Println(nums)

	// var a,b = 0,0

	// for i := 0; i <=(len(nums));i++{
	// 	if a <=nums[i]{
	// 		b=a
	// 		a = nums[i]
	// 	}
	// 	if a > nums[i] && nums[i] > b{
	// 		nums[i] = b
	// 	}
	// }
	// 	fmt.Println(b)

	var arr [10] int
  	for x :=0; x<len(arr); x++ {
    	arr[x]=rand.Int()%99+1
	}
  fmt.Println(arr)
  for x := 0; x<len(arr)/2; x++{
    i := arr[x]
    arr[x]=arr[len(arr)-1-x]
    arr[len(arr)-1-x] = i
  }
  fmt.Println(arr)




}


