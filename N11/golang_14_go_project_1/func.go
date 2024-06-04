package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type Company struct {
  Id int `json:"id"`
  Name    string `json:"name"`
  Manager      string    `json:"manager"`
  Rank 		float32 
  Phone 	string
  Gmail 	string
}

func main() {

  user := Company{
    Id: 1,
    Name:    "John",
    Manager:      "Mary",
	Rank:  2.6,
	Phone: "+998944444444",
	Gmail: "someone@gmail.com",
  }
  file, err := os.OpenFile("Company.json",os.O_CREATE | os.O_APPEND,0644)
  if err != nil {
    fmt.Println("Error", err)
    return
  }
  defer file.Close()
  encoder := json.NewEncoder(file)

  err = encoder.Encode(user)
  if err != nil {
    fmt.Println("Xatolik ", err)
    return
  }

  for i := 0; i< user.Id;i++{
	fmt.Println(user.Gmail)
  }
  
  fmt.Println("Malumotlar JSON faylga muvaffaqiyatli yozildi.")
}















