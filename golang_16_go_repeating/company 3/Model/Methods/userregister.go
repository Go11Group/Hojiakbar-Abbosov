package methods

import (
	"encoding/json"
	"fmt"
	"os"
)

func RegisterUser() {
	users := []User{}
	user := User{}
	// var id, age int
	// var name, surname string
	fmt.Printf("Id kiriting: ")
	fmt.Scan(&user.Id)
	fmt.Printf("Ismingizni kiriting: ")
	fmt.Scan(&user.Name)
	fmt.Printf("Familyangizni kiriting: ")
	fmt.Scan(&user.Surname)
	fmt.Printf("Yoshingizni kiriting: ")
	fmt.Scan(&user.Age)
	f, err := os.OpenFile("users.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	info := json.NewDecoder(f)
	err = info.Decode(&users)
	if err != nil {
		fmt.Println()
	}
	users = append(users, user)

	f.Truncate(0)
	f.Seek(0, 0)

	data := json.NewEncoder(f)
	err = data.Encode(&users)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}


func LoginUser( id int ) bool{
	users := []User{}
	f, err := os.OpenFile("users.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	info := json.NewDecoder(f)
	err = info.Decode(&users)
	if err != nil {
		fmt.Println()
	}
	for i := 0; i < len(users); i++ {
		if  users[i].Id==id{
			return true
		 }
	}
	return false
}