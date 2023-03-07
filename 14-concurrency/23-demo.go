package main

import (
	"fmt"
	"time"
)

func main() {
	fibCh := genFib()
	for no := range fibCh {
		fmt.Println(no)
	}
}

func genFib() <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
		timeoutCh := time.After(5 * time.Second)
	LOOP:
		for {
			select {
			case <-timeoutCh:
				break LOOP
			default:
				ch <- x
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

/*
func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
*/
