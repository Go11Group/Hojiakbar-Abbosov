package main

import (
	"fmt"
)

func main() {

	arr := []byte{'*', '*', '*', '*', '*', '*', '*', '*', '*'}

	isTrue := true

	cas := [][]uint8{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}

	for step := 0; step < 9; step++ {

		for i := 0; i < 9; i++ {
			if (i+1)%3 == 0 {
				fmt.Printf("%c    \n\n", arr[i])
			} else {
				fmt.Printf("%c     ", arr[i])
			}
		}

		var input uint8

		if isTrue {
			isTrue = false
		} else {
			isTrue = true
		}
		if isTrue {
			fmt.Println("2-user")
		} else {
			fmt.Println("1-user")
		}

		for {
			fmt.Println("Enter number 1...9")
			fmt.Scan(&input)
			if arr[input-1] == '*' {
				break
			}
			fmt.Println("False")
		}

		if !isTrue {
			arr[input-1] = 'X'
		} else {
			arr[input-1] = '0'
		}

		for i := 0; i < len(cas); i++ {
			if string(arr[cas[i][0]])+string(arr[cas[i][1]])+string(arr[cas[i][2]]) == "XXX" || string(arr[cas[i][0]])+string(arr[cas[i][1]])+string(arr[cas[i][2]]) == "000" {
				if isTrue {
					fmt.Println("2-user won!!!")
					return
				} else {
					fmt.Println("1-user won!!!")
					return
				}
			}
		}
	}
	fmt.Println("Draw!!!")

}
