package main

import (
  "fmt"
  "time"
)

func IsPrime(num int)(int,string) {
    for i := 2; i <= num; i++ {
        if num%i == 0 {
            return num,"Tub emas"
        }
    }
    return num,"Tub"
}

func main() {
    for i := 2; i < 10; i++ {
        a,b := IsPrime(i)
        go fmt.Println(a,b)
    }

  time.Sleep(5 * time.Second)
}
