package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	ch := generateNos()
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

//producer
func generateNos() <-chan int {
	ch := make(chan int)
	go func() {
		timeoutCh := time.After(20 * time.Second)
		var i = 1
	LOOP:
		for {
			select {
			case ch <- i * 10:
				time.Sleep(500 * time.Millisecond)
				i++
			case <-timeoutCh:
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}

/* func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
} */
