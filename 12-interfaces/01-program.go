package main

import (
	"fmt"
	"math"
)

type Circle struct /* implements AreaFinder */ {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct /* implements AreaFinder */ {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

/* utility functions */
type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

func main() {
	c := Circle{12}
	// fmt.Println("Area :", c.Area())
	PrintArea(c)

	r := Rectangle{10, 20}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)
}
