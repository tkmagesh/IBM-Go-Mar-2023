package main

import (
	"fmt"

	"github.com/admpub/timeago"
	"github.com/fatih/color"
	"github.com/tkmagesh77/ibm-go-mar-2023/09-modularity/calculator"
	"github.com/tkmagesh77/ibm-go-mar-2023/09-modularity/calculator/utils"
)

func init() {
	timeago.Set("language", "en")
	// timeago.Set("location", "America/New_York")
}

func main() {
	// fmt.Println("App executed")
	result := timeago.Take("2019-10-23 10:46:00")
	// fmt.Printf("%T\n", result)
	fmt.Println(result)
	color.Red("App executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Operation Count :", calculator.OpCount())
	fmt.Println("is 21 Even number ?:", utils.IsEven(21))
	Greet("Magesh")
}
