package methods

import (
	"encoding/json"
	"fmt"
	"os"

)





func RegisterManager() {
	f, err := os.OpenFile("manager.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	managers := []Recruiter{}

	manager:=Recruiter{}

	fmt.Printf("Id ni kiriting: ")
	fmt.Scan(&manager.Id)
	fmt.Printf("Ismingizni kiriting: ")
	fmt.Scan(&manager.Name)
	fmt.Printf("Familyangizni kiriting: ")
	fmt.Scan(&manager.Surname)
	fmt.Printf("Password kiriting: ")
	fmt.Scan(&manager.Password)
	fmt.Printf("Email kiriting: ")
	fmt.Scan(&manager.Email)


	info := json.NewDecoder(f)
	err = info.Decode(&managers)
	if err != nil {
		fmt.Println()
	}
	managers = append(managers, manager)

	f.Truncate(0)
	f.Seek(0, 0)

	data := json.NewEncoder(f)
	err = data.Encode(managers)
	if err != nil {
		fmt.Println()
	}

}

func LoginManager(email string, password string) int   {
	f, err := os.OpenFile("manager.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	managers := []Recruiter{}
	data1 := json.NewDecoder(f)
	err = data1.Decode(&managers)
	if err != nil {
		fmt.Println()
	}
	if err != nil {
		fmt.Println()
	}
	for i := 0; i < len(managers); i++ {
		if managers[i].Email == email && managers[i].Password == password {
			return managers[i].Id
			
		}
	}
	return -1
	

}
