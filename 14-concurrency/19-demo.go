package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		/*
			go func() {
				defer wg.Done()
				fmt.Println(i)
			}()
		*/
		go func(no int) {
			defer wg.Done()
			fmt.Println(no)
		}(i)
	}
	wg.Wait()
}
