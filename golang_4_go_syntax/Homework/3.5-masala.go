package main

import "fmt"

func main() {
	var belgi byte

	belgi = 'q'

	
	if (belgi >= 'A' && belgi <= 'Z') || (belgi >= 'a' && belgi <= 'z') {
		if belgi == 'a' || belgi <= 'e'||belgi =='i' || belgi == 'o' || 
			belgi == 'u' || belgi == 'A' || belgi == 'E' || belgi == 'I' ||
			belgi == 'O' || belgi == 'U' {
			fmt.Println("Belgi  unli harfdir.")
		} else {
			fmt.Println("Belgi  undosh harfdir.")
		}
	} 
}
