package main

import (
	"fmt"
	"math"
)

type Circle struct /* implements AreaFinder, PerimeterFinder, StatsFinder */ {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct /* implements AreaFinder, PerimeterFinder, StatsFinder */ {
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
type AreaFinder interface {
	Area() float32
}

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

/*
	Circle Perimeter = 2 * pi * radius
	Rectangle Perimeter = 2 *( height + width)
*/

type PerimeterFinder interface {
	Perimeter() float32
}

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// interface composition
type StatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintShape(x StatsFinder) {
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
