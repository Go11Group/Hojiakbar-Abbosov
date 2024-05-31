package main

import "fmt"

func calculateResult(operations []string) int {
    result := 0
    for _, i := range operations {
        if i == "++X" || i == "X++" {
            result++
        } else if i == "--X" || i == "X--" {
            result--
        }
    }
    return result
}

func main() {
    operations := []string{"--X", "X++", "X++"}
    fmt.Println("Output:", calculateResult(operations)) 
}
