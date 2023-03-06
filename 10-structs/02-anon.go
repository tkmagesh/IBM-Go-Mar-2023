package main

import "fmt"

func main() {
	product := struct {
		Id   int
		Name string
		Cost float32
	}{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(product.Id, product.Name, product.Cost)
	fmt.Println(Format(product))

	zeroByteObj := struct{}{}
	fmt.Println(zeroByteObj)
}

func Format(p struct {
	Id   int
	Name string
	Cost float32
}) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}
