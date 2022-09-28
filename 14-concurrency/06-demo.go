package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	flag.IntVar(&count, "count", 0, "the number of goroutines to spawn")
	flag.Parse()

	rand.Seed(700)
	wg := &sync.WaitGroup{}

	fmt.Println("Hit ENTER to start....")
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1)    // wg counter = 1
		go fn(i, wg) //schedule the execution
	}

	wg.Wait() //=> block until the wg counter becomes 0
	fmt.Println("Done")
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() //=> decrementing the wg counter by 1
	t := rand.Intn(10000)
	fmt.Printf("fn[%d] started (%d)\n", id, t)
	time.Sleep(time.Duration(t) * time.Millisecond)
	fmt.Printf("fn[%d] completed\n", id)
}
