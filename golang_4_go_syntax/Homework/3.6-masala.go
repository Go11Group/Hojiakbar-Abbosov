package main

import "fmt"

func main() {
	var belgi byte

	belgi = 'Q'

	
	if belgi >= 'A' && belgi <= 'Z' {
		fmt.Println("Belgi  katta harfdir.")
	}else if belgi >= 'a' && belgi <= 'z' {
		fmt.Println("Belgi  kichik harfdir.")
	}
} 
