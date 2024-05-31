package methods

import (
	"encoding/json"
	"fmt"
	"os"
)

type Meniger struct {
	Password string `json:"password"`
	Gmail    string `json:"gmail"`
}

func Login(p, g string) {

	m := []Meniger{}
	m1 := Meniger{}
	m1.Gmail = g
	m1.Password = p

	f, err := os.OpenFile("login.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	info := json.NewDecoder(f)
	err = info.Decode(&m)
	if err != nil {
		fmt.Println()
	}
	m = append(m, m1)

	f.Truncate(0)
	f.Seek(0, 0)

	data := json.NewEncoder(f)
	err = data.Encode(&m)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}
