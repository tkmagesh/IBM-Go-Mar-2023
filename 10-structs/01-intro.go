package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

func main() {
	// var product Product
	// var product Product = Product{100, "Pen", 10}
	// var product Product = Product{Id: 100, Name: "Pen", Cost: 10}
	/*
		var product Product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/
	/*
		var product = Product{
			Id:   100,
			Name: "Pen",
			Cost: 10,
		}
	*/
	product := Product{
		Id:   100,
		Name: "Pen",
		Cost: 10,
	}
	// fmt.Println(product)
	// fmt.Printf("%#v\n", product)
	fmt.Printf("%+v\n", product)

	fmt.Println("Accessing the fields")
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", product.Id, product.Name, product.Cost)

	fmt.Println("Updating the fields")
	product.Cost = 9
	fmt.Printf("Id = %d, Name = %q, Cost = %0.2f\n", product.Id, product.Name, product.Cost)

	fmt.Println("Structs are values")
	newProduct := product // assignment will result in a new copy
	newProduct.Name = "Fountain Pen"
	fmt.Printf("product : Id = %d, Name = %q, Cost = %0.2f\n", product.Id, product.Name, product.Cost)
	fmt.Printf("newProduct : Id = %d, Name = %q, Cost = %0.2f\n", newProduct.Id, newProduct.Name, newProduct.Cost)

	fmt.Println("Equality")
	p2 := Product{Id: 100, Name: "Pen", Cost: 9}
	fmt.Println("product : ", product)
	fmt.Println("p2 	 : ", p2)
	fmt.Println("product == p2 :", product == p2) // struct instances are compared by value

	fmt.Println("Pointers")
	/*
		pencilValue := Product{Id: 200, Name: "Pencil", Cost: 5}
		pencil := &pencilValue
	*/
	pencil := &Product{Id: 200, Name: "Pencil", Cost: 5}
	fmt.Println((*pencil).Id)
	fmt.Println(pencil.Id) // fields can be accessed WITHOUT dereferencing the pointer of a struct
	fmt.Printf("%p\n", pencil)

	/*
		pencilValue := Product{Id: 200, Name: "Pencil", Cost: 5}
		pencilPtr1 := &pencilValue
		pencilPtr2 := &pencilValue

		pencilPtr1.Cost = 4
		fmt.Println(pencilValue)
		fmt.Println(pencilPtr1)
		fmt.Println(pencilPtr2)
	*/

	fmt.Println("Before applying discount")
	fmt.Println(Format(product))
	ApplyDiscount(&product, 10)
	fmt.Println("After applying discount")
	fmt.Println(Format(product))
}

/*
	Write the following functions
		Format(?)
			return the formatted string (below) for the given product
				Id = <val>, Name = "<val>", Cost = <val>
		ApplyDiscount(?)
			update the given product cost after applying the given discount percentage
*/

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %0.2f", p.Id, p.Name, p.Cost)
}

func ApplyDiscount(p *Product, discountPercentage float32) {
	p.Cost = p.Cost * ((100 - discountPercentage) / 100)
}
