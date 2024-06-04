package main

import (
	"fmt"
	"my_module/Homework"
)

func main() {
	
	rect := Homework.Rectangle{Width: 5, Height: 6}
	c := Homework.Circle{Radius: 7}

	fmt.Printf("Rectangle Area: %f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %f\n", rect.Perimeter())
	fmt.Printf("Rectangle Diagonal: %f\n", rect.Diagonal())

	fmt.Printf("Circle Area: %f\n", c.Area())
	fmt.Printf("Circle Perimeter: %f\n", c.Perimeter())
	fmt.Printf("Circle Diagonal: %f\n", c.Diagonal())
}