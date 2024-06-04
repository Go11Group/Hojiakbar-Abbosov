package main

import "fmt"

func swap(a, b, c *int) {
    *a, *b, *c = *c, *a, *b
}

func main() {
    x, y, z := 1, 2, 3
    fmt.Println("Old values:", x, y, z) 

    swap(&x, &y, &z)
    fmt.Println("New values:", x, y, z) 
}
