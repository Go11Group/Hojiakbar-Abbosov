package main

import (
 "fmt"
)

func main() {
 	arr := []byte{'*', '*', '*', '*', '*', '*', '*', '*', '*'}
 	playGame(arr)
}

func playGame(arr []byte) {
 	cas := [][]uint8{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {0, 3, 6}, {1, 4, 7}, {2, 5, 8}, {0, 4, 8}, {2, 4, 6}}
 	isTrue := true

 	for step := 0; step < 9; step++ {
  		printBoard(arr)
  		var input uint8

	if isTrue {
		fmt.Println("1-user")
	} else {
		fmt.Println("2-user")
	}

	for {
		fmt.Println("Enter a number (1-9):")
		fmt.Scan(&input)
		if input >= 1 && input <= 9 && arr[input-1] == '*' {
			break
	}
		fmt.Println("Invalid input. Please try again.")
	}

	if !isTrue {
		arr[input-1] = 'X'
	} else {
		arr[input-1] = 'O'
	}

	if checkWin(arr, cas) {
		printBoard(arr)
	if isTrue {
		fmt.Println("1-user won!!!")
	} else {
		fmt.Println("2-user won!!!")
	}
	return
	}

	isTrue = !isTrue
	}

	printBoard(arr)
	fmt.Println("Draw!!!")
	}

	func printBoard(arr []byte) {
		for i := 0; i < 9; i++ {
			if (i+1)%3 == 0 {
				fmt.Printf("%c    \n\n", arr[i])
			} else {
				fmt.Printf("%c     ", arr[i])
			}
		}
	}

	func checkWin(arr []byte, cas [][]uint8) bool {
		for i := 0; i < len(cas); i++ {
			if arr[cas[i][0]] == arr[cas[i][1]] && arr[cas[i][1]] == arr[cas[i][2]] && arr[cas[i][0]] != '*' {
				return true
			}
		}
			return false
}