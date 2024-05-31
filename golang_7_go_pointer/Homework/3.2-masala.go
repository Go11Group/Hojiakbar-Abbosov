package main

import (
	"fmt"
	"strings"
)

func preserveDigits(str string) {
	var digits strings.Builder

	for _, char := range str {
		if char >= '0' && char <= '9' {
			digits.WriteRune(char)
		}
	}

	newStr := digits.String()
	fmt.Println(newStr)
}

func main() {
	inputStr := "abc123def456ghi789"
	preserveDigits(inputStr) 
}

