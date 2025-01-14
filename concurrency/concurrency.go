package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go Routines")

	go func() {
		defer wg.Done()

		for char := 'a'; char < 'a'+26; char++ {

			time.Sleep(200 * time.Millisecond)

			fmt.Printf("%c ", char)
		}
		println("")

	}()

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("%d ", number)
		}
		println("")
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\n Terminating Program")
}
