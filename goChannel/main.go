package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// data := <- a // read from channel
// a <- data // write to channel
func main() {
	// var a chan int
	a := make(chan int)
	if a == nil {
		fmt.Println("channel a is nil")
		fmt.Printf("Type of a is %T\n", a)
	}

	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				result := doWork()
				a <- result
			}()

			// result := doWork()
			// a <- result
		}
		wg.Wait()
		close(a)

	}()
	for data := range a {
		fmt.Println(data)
	}
}

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
