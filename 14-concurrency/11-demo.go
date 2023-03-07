package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	/* Use case - 1 */
	/*
		ch <- 100
		data := <- ch
		fmt.Println(data)
	*/

	/* Use Case - 2 */
	/*
		data := <- ch
		ch <- 100
		fmt.Println(data)
	*/

	/* Use Case - 3 */
	/*
		ch <- 100
		go func() {
			data := <-ch
			fmt.Println(data)
		}()
	*/

	/* Use Case - 4 */
	/*
		go func() {
			ch <- 100
		}()
		data := <-ch
		fmt.Println(data)
	*/

	/* Use Case - 5 */
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		data := <-ch
		fmt.Println(data)
		wg.Done()
	}()
	ch <- 100
	wg.Wait()
}
