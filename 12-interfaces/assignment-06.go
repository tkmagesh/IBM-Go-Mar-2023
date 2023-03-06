package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Write the apis for the following

IndexOf => return the index of the given product (return -1 if not exists )
	ex:  returns the index of the given product

Includes => return true if the given product is present in the collection else return false
	ex:  returns true if the given product is present in the collection

Filter => return a new collection of products that satisfy the given condition
	some of the use cases:
		1. filter all costly products (cost > 1000)
			OR
		2. filter all stationary products (category = "Stationary")
			OR
		3. filter all costly stationary products
		etc

Any => return true if any of the product in the collections satifies the given criteria
	some of the use cases:
		1. are there any costly product (cost > 1000)?
		OR
		2. are there any stationary product (category = "Stationary")?
		OR
		3. are there any costly stationary product?
		etc

All => return true if all the products in the collections satifies the given criteria
	some of use cases:
		1. are all the products costly products (cost > 1000)?
		OR
		2. are all the products stationary products (category = "Stationary")?
		OR
		3. are all the products costly stationary products?
		etc
*/

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Units = %d, Category = %q", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

type Products []Product

func (products Products) String() string {
	var sb strings.Builder
	for _, product := range products {
		sb.WriteString(fmt.Sprintf("%s\n", product))
	}
	return sb.String()
}

func (products Products) IndexOf(p Product) int {
	for idx, product := range products {
		if product == p {
			return idx
		}
	}
	return -1
}

/*
func (products Products) FilterCostlyProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Cost > 1000 {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) FilterStationaryProducts() Products {
	result := Products{}
	for _, product := range products {
		if product.Category == "Stationary" {
			result = append(result, product)
		}
	}
	return result
}
*/

func (products Products) Filter(predicate func(Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

func (products Products) All(predicate func(Product) bool) bool {
	for _, product := range products {
		if !predicate(product) {
			return false
		}
	}
	return true
}

func CostlyProductPredicate(cost float32) func(Product) bool {
	return func(p Product) bool {
		return p.Cost > cost
	}
}

/*
Sort => to sort the products by any attribute
	IMPORTANT: use sort.Sort() function
*/

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(byName{products})
	case "Cost":
		sort.Sort(byCost{products})
	case "Units":
		sort.Sort(byUnits{products})
	case "Category":
		sort.Sort(byCategory{products})
	}
}

func (products Products) Sort2(attrName string) {
	switch attrName {
	case "Id":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Id < products[j].Id
		})
	case "Name":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Name < products[j].Name
		})
	case "Cost":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Cost < products[j].Cost
		})
	case "Units":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Units < products[j].Units
		})
	case "Category":
		sort.Slice(products, func(i, j int) bool {
			return products[i].Category < products[j].Category
		})
	}
}

// sort.Interface implementation
func (products Products) Len() int {
	return len(products)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

// to sort by name
type byName struct {
	Products
}

func (byName byName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

// to sort by Cost
type byCost struct {
	Products
}

func (byCost byCost) Less(i, j int) bool {
	return byCost.Products[i].Cost < byCost.Products[j].Cost
}

// to sort by Units
type byUnits struct {
	Products
}

func (byUnits byUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

// to sort by Category
type byCategory struct {
	Products
}

func (byCategory byCategory) Less(i, j int) bool {
	return byCategory.Products[i].Category < byCategory.Products[j].Category
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
	stove := Product{102, "Stove", 5000, 5, "Utencil"}
	fmt.Println("Index of stove :", products.IndexOf(stove))

	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Costly Products")
	/*
		costlyProductPredicate := func(p Product) bool {
			return p.Cost > 1000
		}
	*/
	// costlyProducts := products.FilterCostlyProducts()
	costlyProductPredicate := CostlyProductPredicate(2000)
	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println(costlyProducts)

	fmt.Println("Stationary Products")
	// stationaryProducts := products.FilterStationaryProducts()
	stationaryProductPredicate := func(p Product) bool {
		return p.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println(stationaryProducts)

	fmt.Println("Are all the products costly products ? :", products.All(costlyProductPredicate))
	fmt.Println("Are there any costly products ? :", products.Any(costlyProductPredicate))

	fmt.Println("Default Sort")
	// sort.Sort(products)
	// products.Sort("Id")
	products.Sort2("Id")
	fmt.Println(products)

	fmt.Println("Sort By Name")
	// sort.Sort(byName{products})
	// products.Sort("Name")
	products.Sort2("Name")
	fmt.Println(products)

	fmt.Println("Sort By Cost")
	// sort.Sort(byCost{products})
	// products.Sort("Cost")
	products.Sort2("Cost")
	fmt.Println(products)

	fmt.Println("Sort By Units")
	// sort.Sort(byUnits{products})
	// products.Sort("Units")
	products.Sort2("Units")
	fmt.Println(products)

	fmt.Println("Sort By Category")
	// sort.Sort(byCategory{products})
	// products.Sort("Category")
	products.Sort2("Category")
	fmt.Println(products)
}
