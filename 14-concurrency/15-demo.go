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
	go func() {
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 10
		}
	}()
	return ch
}
