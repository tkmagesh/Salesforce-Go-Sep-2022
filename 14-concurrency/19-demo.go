package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()

	go func() {
		time.Sleep(3 * time.Second)
		no := <-ch3
		fmt.Println(no)
	}()

	/*
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	*/
	for i := 0; i < 3; i++ {
		select {
		case n1 := <-ch1:
			fmt.Println(n1)
		case n2 := <-ch2:
			fmt.Println(n2)
		case ch3 <- 300:
			fmt.Println("Sent 300 to ch3")
		}
	}
}
