package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Channels() {
	rand.Seed(time.Now().UnixNano())
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok {
			fmt.Printf("Random 2 digit integer: %d\n", msg)
			wg.Done()
			return
		}

		fmt.Println("Channel closed")
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		min := 10
		max := 99
		ch <- rand.Intn(max-min+1) + min
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
