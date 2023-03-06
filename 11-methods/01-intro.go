package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) WhoAmI() {
	fmt.Println("I am a product :", p.Name)
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

func main() {
	pen := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	fmt.Println(pen.Format())
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())
}
