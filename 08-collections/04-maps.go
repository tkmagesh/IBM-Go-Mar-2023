package main

import "fmt"

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	// var productRanks = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	// productRanks := map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	/*
		productRanks := map[string]int{
			"pen":    5,
			"pencil": 1,
			"marker": 3,
		}
	*/
	// productRanks := make(map[string]int)
	productRanks := map[string]int{}
	productRanks["pen"] = 5
	productRanks["pencil"] = 1
	productRanks["marker"] = 3

	fmt.Println(productRanks)
	fmt.Println("len(productRanks) :", len(productRanks))

	fmt.Println("Appending a new key-value pair")
	productRanks["notebook"] = 4
	fmt.Println(productRanks)

	fmt.Println("Iterating a map")
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// keyToCheck := "pencil"
	keyToCheck := "scribble-pad"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Product %q does not exist\n", keyToCheck)
	}

	keyToDelete := "scribble-pad"
	delete(productRanks, keyToDelete)
	fmt.Printf("After deleting key - %q\n", keyToDelete)
	fmt.Println(productRanks)
}
