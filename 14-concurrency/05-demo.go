package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1) // wg counter = 1
		go f1(wg) //schedule the execution
	}
	f2()
	wg.Wait() //=> block until the wg counter becomes 0

}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() //=> decrementing the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
