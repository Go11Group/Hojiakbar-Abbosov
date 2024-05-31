package main

import (
 "encoding/json"
 "fmt"
 "os"
)

type User struct {
 Username string 
 Email    string 
 Age      int    
}

func main() {
 
 file, err := os.OpenFile("malumotlar.json", os.O_RDWR|os.O_CREATE, 0644)
 if err != nil {
  fmt.Println("Xatolik", err)
  return
 }
 defer file.Close()

 var existingData User
 err = json.NewDecoder(file).Decode(&existingData)
 if err != nil  {
  fmt.Println("Xatolik", err)
  return
 }

 newUser := User{
  Username: "Abdulloh",
  Email:    "abdulloh@example.com",
  Age:      25,
 }

 encoder := json.NewEncoder(file)
 if err != nil {
  fmt.Println("Xatolik", err)
  return
 }

 err = encoder.Encode(newUser)
 if err != nil {
  fmt.Println("Xatolik", err)
  return
 }

 fmt.Println("Yangi ma'lumotlar JSON faylga muvaffaqiyatli yozildi.")
}