package methods

import (
	"encoding/json"
	"fmt"
	"os"
)




func RegisterManager() {
	managers := []Recruiter{}
	f, err := os.OpenFile("manager.json", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	var id int
	var name, surname, email, password string

	fmt.Printf("Id ni kiriting: ")
	fmt.Scan(&id)
	fmt.Printf("Ismingizni kiriting: ")
	fmt.Scan(&name)
	fmt.Printf("Familyangizni kiriting: ")
	fmt.Scan(&surname)
	fmt.Printf("Password kiriting: ")
	fmt.Scan(&password)
	fmt.Printf("Email kiriting: ")
	fmt.Scan(&email)

	manager := Recruiter{
		Name: name,
		Id: id,
		Password: password,
		Email: email,
		Surname: surname,
	}
	// managers = append(managers, manager)
	
	
	defer f.Close()

	info := json.NewDecoder(f)
	err = info.Decode(&managers)
	if err != nil {
		fmt.Println()
	}
	managers = append(managers, manager)

	f.Truncate(0)
	f.Seek(0, 0)

	data := json.NewEncoder(f)
	err = data.Encode(&managers)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	
}

