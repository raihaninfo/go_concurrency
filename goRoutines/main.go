package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main Starts")
	go hello(1)
	go hello(2)
	go hello(3)
	go hello(4)
	go hello(5)
	fmt.Println("Main End")

	time.Sleep(100 * time.Millisecond)
}

func hello(grn int) {
	for i := 0; i < 3; i++ {
		fmt.Println(grn, ": Hello")
	}
}
