package main

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/tkmagesh77/ibm-go-mar-2023/09-modularity/calculator"
	"github.com/tkmagesh77/ibm-go-mar-2023/09-modularity/calculator/utils"
)

func main() {
	// fmt.Println("App executed")
	color.Red("App executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Operation Count :", calculator.OpCount())
	fmt.Println("is 21 Even number ?:", utils.IsEven(21))
	Greet("Magesh")
}
