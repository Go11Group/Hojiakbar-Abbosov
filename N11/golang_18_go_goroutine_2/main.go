package main

import (
  "fmt"
  "sync"
)

func main() {
  bug := 0
  for i := 0; i < 1000000; i++ {
    sum := 0
    w := sync.WaitGroup{}
    w.Add(10)
    for i := 1; i <= 1000; i += 100 {
      go func(num int) {
        defer w.Done()
        localSum := 0
        for j := 0; j < 100; j++ {
          localSum += num + j
        }
        sum += localSum
      }(i)
    }
    w.Wait()
    if sum == 500500 {
      bug++
    }
  }
  fmt.Println(bug)

}
