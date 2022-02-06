package main

import "fmt"

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
		for i := 0; i < 100; i++ {
			a <- i // write to channel
		}
		close(a)
	}()

	for data := range a {
		fmt.Println(data)
	}
}
