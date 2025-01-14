package main

import (
	"log"
	"time"
)

func main() {

	ticker := time.NewTicker(1 * time.Second)

	go run(ticker)

	n := 0
	for {
		time.Sleep(2 * time.Second)
		n += 1
		println(n)

		if n == 3 {
			ticker.Stop()
			println("stop!")
		}
	}

}

func run(ticker *time.Ticker) {
	log.Println("hello")

	for _ = range ticker.C {
		// simply push blank message
		println("in")
	}

	println("&&&&&&&&&")

	return
}
