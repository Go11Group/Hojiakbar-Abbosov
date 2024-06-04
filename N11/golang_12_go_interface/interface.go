package main

import "fmt"

type geometry interface {
 area() float32
 perimetr() float32
}

type rectange struct {
 width, height float32
}

func (r rectange) area() float32 {
 return r.width * r.height
}

func (r rectange) perimetr() float32 {
 return 2 * (r.height + r.width)
}

func (r rectange) dioganal() float32 {
 return 0
}

type circle struct {
 radius float32
}

func (c circle) area() float32 {
 return c.radius * c.radius * 3.14
}

func (c circle) perimetr() float32 {
 return 2 * 3.14 * c.radius
}

func printer(rec geometry) {
 fmt.Printf("Area: %f\nPerimeter:%f\n", rec.area(), rec.perimetr())
}

func main() {
 rectan := rectange{5, 6}
 c := circle{7}

 printer(rectan)
 printer(c)


}
