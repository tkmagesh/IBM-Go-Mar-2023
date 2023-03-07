package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                 //=> 30
	fmt.Println(sum(10, 20, 30, 40))                         //=> 100
	fmt.Println(sum(10))                                     //=> 10
	fmt.Println(sum())                                       //=> 0
	fmt.Println(sum(10, "20", 30, "40"))                     //=> 100 (use strconv.Atoi())
	fmt.Println(sum(10, "20", 30, "40", "abc"))              //=> 100
	fmt.Println(sum(10, 20, []int{30, 40}))                  //=> 100
	fmt.Println(sum(10, 20, []any{30, 40, []int{10, 20}}))   //=> 130
	fmt.Println(sum(10, 20, []any{30, 40, []any{10, "20"}})) //=> 130
}

func sum(values ...any) int {
	result := 0
	for _, val := range values {
		switch v := val.(type) {
		case int:
			result += v
		case string:
			if x, err := strconv.Atoi(v); err == nil {
				result += x
			}
		case []any:
			result += sum(v...)
		case []int:
			var list []any = make([]any, 0, len(v))
			for _, x := range v {
				list = append(list, x)
			}
			result += sum(list...)
		}
	}
	return result
}
