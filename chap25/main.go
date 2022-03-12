package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate: //Context
			fmt.Println("Terminate")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
	wg.Done()
}

var wg2 sync.WaitGroup

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)
	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}
