package main

import "fmt"

func main() {
	var belgi byte

	belgi = '1'

	if belgi >= '0' && belgi <= '9' {
		fmt.Println("Belgi raqamdir.")
	}else if (belgi >= 'A' && belgi <= 'Z') || (belgi >= 'a' && belgi <= 'z') {
		fmt.Println("Belgi harfdir")
	}else {
		fmt.Println("Belgi simvol.")
	}
}
