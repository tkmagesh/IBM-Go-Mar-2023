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

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct /* implements AreaFinder */ {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

/* utility functions */

func PrintArea(x interface {
	Area() float32
}) {
	fmt.Println("Area :", x.Area())
}

func PrintPerimeter(x interface {
	Perimeter() float32
}) {
	fmt.Println("Perimeter :", x.Perimeter())
}

/* interface composition */

/*
func PrintShape(x interface {
	Area() float32
	Perimeter() float32
}) {
	PrintArea(x)
	PrintPerimeter(x)
}
*/

func PrintShape(x interface {
	interface {
		Area() float32
	}
	interface {
		Perimeter() float32
	}
}) {
	PrintArea(x)
	PrintPerimeter(x)
}

func main() {
	c := Circle{12}
	// fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{10, 20}
	// fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}
