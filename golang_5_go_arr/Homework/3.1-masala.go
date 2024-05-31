package main

import (
	"fmt"
)

func main() {
	Login := "Hoji123"
	Password := "parol123"

	fmt.Print("Loginni kiriting: ")
	var login string
	fmt.Scanln(&login)

	fmt.Print("Parolni kiriting: ")
	var password string
	fmt.Scanln(&password)

	if login == Login && password == Password {
		fmt.Println("Login va parol to'g'ri")
	} else {
		fmt.Println("Login yoki parol noto'g'ri")
	}
}
