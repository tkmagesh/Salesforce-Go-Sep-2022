package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //schedule the execution
	f2()
	time.Sleep(1 * time.Second) //blocking the main func for 1 second
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
