package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10
	ch := generateNos(count)
	for i := 0; i < count; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Done")
}

func generateNos(count int) <-chan int {
	ch := make(chan int)
	for i := 1; i <= count; i++ {
		go func(no int) {
			time.Sleep(500 * time.Millisecond)
			ch <- no * 10
		}(i)
	}
	return ch
}
