package main 

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 float64
	var operator string

	num1 = 5
	num2 = 4
	operator = "*"

	var result float64

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Nolga bo'lish mumkin emas!")
			return
		}
	case "%":
		result = math.Mod(num1, num2)
	case "^":
		result = math.Pow(num1, num2)
	default:
		fmt.Println("Noto'g'ri operatsiya kiritildi.")
		return
	}

	fmt.Printf("Natija: %.2f\n", result)
}

