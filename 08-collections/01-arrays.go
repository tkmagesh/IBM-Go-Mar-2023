package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}
	// nos := [5]int{3, 1, 4, 2, 5}
	// nos := [5]int{3, 1, 4}

	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Accessing elements")
	fmt.Println(nos[0], nos[1], nos[2])

	fmt.Println("Iterating an array using indexer")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an array using range")
	for i, val := range nos {
		fmt.Printf("nos[%d] = %d\n", i, val)
	}
}
