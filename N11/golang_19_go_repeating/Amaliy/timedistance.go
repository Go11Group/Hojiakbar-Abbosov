package main

import (
    "fmt"
    "time"
)

func timeDistance(ch chan time.Time, resultCh chan time.Duration, second time.Time) {
    currentTime := <-ch
    duration := currentTime.Sub(second)
    resultCh <- duration
}

func main() {
    ch := make(chan time.Time)
    resultCh := make(chan time.Duration)
    currentTime := time.Now()
    secondTime := currentTime.Add(time.Hour * 2) 

    go timeDistance(ch, resultCh, secondTime)

    ch <- currentTime

    result := <-resultCh
    fmt.Println("Time difference:", result)
}

