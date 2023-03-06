package main

import "fmt"

type Employee struct {
	Id   int
	Name string
	Org  Organization
}

type Organization struct {
	Name string
	City string
}

/* Modify the code in such a way that when the fields of "ibm" are modified, the changes are reflected in emp1 & emp2 */

func main() {
	ibm := Organization{Name: "IBM", City: "Bengaluru"}
	emp1 := Employee{Id: 100, Name: "Magesh", Org: ibm}
	fmt.Println(emp1.Org)
	emp2 := Employee{Id: 102, Name: "Suresh", Org: ibm}
	fmt.Println(emp2.Org)

	ibm.City = "Pune"
	fmt.Println(emp1.Org)
	fmt.Println(emp2.Org)
}
