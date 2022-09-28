package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	ch := generateNos()
	/*
		for {
			if no, ok := <-ch; !ok {
				break
			} else {
				fmt.Println(no)
			}
		}
	*/
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

//producer
func generateNos() <-chan int {
	count := 10
	ch := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
		close(ch)
	}()
	return ch
}
