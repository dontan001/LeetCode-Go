package main

import (
	"fmt"
	"sync"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i
	}
	close(ch)
}

func consume(ch <-chan int) {
	for v := range ch {
		fmt.Printf("Consuming %d\n", v)
	}
}

func produceAndConsume() {
	ch := make(chan int, 5)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		produce(ch)
	}()

	go func() {
		defer wg.Done()
		consume(ch)
	}()

	wg.Wait()
	fmt.Printf("Done...\n")
}
