package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	done := make(chan bool)
	ch := generateNos(done)
	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		done <- true
	}()
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

//producer
func generateNos(done <-chan bool) <-chan int {
	ch := make(chan int)
	go func() {
		var i = 1
	LOOP:
		for {
			select {
			case ch <- i * 10:
				time.Sleep(500 * time.Millisecond)
				i++
			case <-done:
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
