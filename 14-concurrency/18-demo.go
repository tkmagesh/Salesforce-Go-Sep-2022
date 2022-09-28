package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 10
	ch := generateNos(count)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Done")
}

//producer
func generateNos(count int) <-chan int {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	go func() {
		for i := 1; i <= count; i++ {
			wg.Add(1)
			go produceValue(i, wg, ch)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

func produceValue(no int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	ch <- no * 10
}
