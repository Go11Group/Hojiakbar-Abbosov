package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type Data struct {
  UsuerId int    `json:"userId"`
  Id      int    `json:"id"`
  Titel   string `json:"title"`
  Bady    string `json:"body"`
}

func main() {
  f, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

  if err != nil {
    panic(err)
  }
  defer f.Close()

  // _, err = f.WriteString("\n jknjdnndgnsklsdflksd")
  // if err != nil {
  //   panic(err)
  // }

  f.Seek(0, 0)

  // fmt.Println(n)

  a := make([]byte, 1200)
  _, err = f.Read(a)
  if err != nil {
    panic(err)
  }
  data := []Data{}

  
  err = json.Unmarshal(a, &data)
  if err != nil {
    panic(err)
  }

  fmt.Print(data)
  for i := 0; i < len(data); i++ {
    fmt.Println(data[i].Bady, data[i].Id, data[i].Titel, data[i].UsuerId)
  }

}
