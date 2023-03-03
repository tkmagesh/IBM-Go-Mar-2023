package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("len(nos) :", len(nos))

	fmt.Println("Iterating a slice using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating a slice using range")
	for i, val := range nos {
		fmt.Printf("nos[%d] = %d\n", i, val)
	}

	fmt.Println("appending to a slice")
	nos = append(nos, 10)
	nos = append(nos, 20, 30, 40)
	hundereds := []int{100, 200, 300}
	nos = append(nos, hundereds...)
	fmt.Println(nos)

	fmt.Println("slicing a slice")
	fmt.Println("nos[2:6] =", nos[2:6])
	fmt.Println("nos[2:] =", nos[2:])
	fmt.Println("nos[:6] =", nos[:6])

	fmt.Println("Inserting a value")
	newNos := append([]int{}, nos[:5]...)
	newNos = append(newNos, 999)
	newNos = append(newNos, nos[5:]...)
	fmt.Println(newNos)

}
