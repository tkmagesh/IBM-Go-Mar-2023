package main

import "fmt"

func main() {
	// var x interface{} // before go 1.18
	var x any //from go 1.18

	x = 100
	x = "Ipsum veniam voluptate ullamco aute ex enim et ullamco adipisicing. Ullamco nisi consequat anim consequat. Lorem eu est ex ullamco occaecat anim sit eu ad excepteur do velit irure adipisicing. Et Lorem est duis dolore ex commodo voluptate ullamco. Consectetur et Lorem officia ut aliqua ut laboris deserunt esse."
	x = true
	x = 99.98
	x = struct{}{}
	fmt.Println(x)

	x = "Duis incididunt cillum commodo ipsum eu exercitation officia id cillum officia irure minim enim veniam. Anim consectetur officia in dolore laborum magna laboris cupidatat aliqua exercitation cupidatat magna. Nostrud eiusmod excepteur deserunt dolor sit culpa. Eiusmod labore incididunt mollit eiusmod ex tempor labore aliquip irure fugiat do quis duis fugiat."
	// x = 100
	// fmt.Println(x.(int) + 200)
	if val, ok := x.(int); ok {
		fmt.Println(val + 200)
	} else {
		fmt.Println("x is not an int")
	}

	// x = 100
	// x = "Veniam laboris veniam ipsum duis ex commodo irure dolor enim aliqua culpa. Anim minim laboris elit proident cillum fugiat. Qui nostrud nisi nulla veniam nisi nulla. Et qui enim duis labore pariatur incididunt. Ea irure aute voluptate consequat eiusmod veniam laboris."
	// x = true
	// x = 99.99
	// x = struct{}{}
	x = struct {
		Id int
	}{Id: 100}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case float64:
		fmt.Println("x is a float64, 10% of x =", val*0.1)
	case struct{}:
		fmt.Println("x is a struct{}")
	default:
		fmt.Println("unknown type")
	}
}
