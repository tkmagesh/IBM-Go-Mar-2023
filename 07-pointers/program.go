package main

import "fmt"

func main() {

	var no int
	no = 100

	var noPtr *int
	noPtr = &no //value -> address

	fmt.Println(no, noPtr)

	//deferencing (address -> value)
	fmt.Println(*noPtr)

	fmt.Println(no == *(&no))

	fmt.Println("Before incrementing, no :", no)
	increment(&no)
	fmt.Println("After incrementing, no :", no)

	n1, n2 := 100, 200
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap( /*  */ )
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n2, n2)
}

func increment(x *int) {
	fmt.Println("x =", x)
	*x++
}

func swap( /* ? */ ) {
	/* ? */
}
