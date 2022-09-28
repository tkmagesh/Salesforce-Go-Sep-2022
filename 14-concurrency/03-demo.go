package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //schedule the execution
	f2()
	//time.Sleep(1 * time.Second) //blocking the main func for 1 second
	fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
