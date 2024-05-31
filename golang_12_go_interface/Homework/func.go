package Homework

import "math"

type Geometry interface {
	Area() float64
	Perimeter() float64
	Diagonal() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Diagonal() float64 {
	return math.Sqrt(r.Width*r.Width + r.Height*r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Diagonal() float64 {
	return 0 
}
