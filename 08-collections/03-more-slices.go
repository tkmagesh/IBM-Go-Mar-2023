package main

import "fmt"

func main() {
	// nos := []int{}
	// nos := make([]int, 0, 3)
	nos := make([]int, 2, 3)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 3)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 5)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 2)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 1)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 4)
	fmt.Printf("len(nos) : %d, cap(nos) : %d, nos : %v\n", len(nos), cap(nos), nos)

	// newNos := nos[2:4]
	newNos := make([]int, 2)
	copy(newNos, nos[2:4])
	newNos[0] = 100
	fmt.Println("nos :", nos)
	fmt.Println("newNos :", newNos)
}
