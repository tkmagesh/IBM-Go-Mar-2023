package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func (p *Product) ApplyDiscount(discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

// Overriding the Product.Format() method
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry : %q", pp.Product.Format(), pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {
	grapes := NewPerishableProduct(100, "Grapes", 50, "2 Days")
	fmt.Println(grapes.Format())
	grapes.ApplyDiscount(10)
	fmt.Println(grapes.Format())
}
