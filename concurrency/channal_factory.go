package main

import (
	"time"
)

func main() {
	go suck(pump())
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func suck(ch chan int) {
	for {
		println(<-ch)
	}
}
