package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main Starts")
	go hello(1)
	fmt.Println("Main End")

	time.Sleep(100 * time.Millisecond)
}

func hello(grn int) {
	for i := 0; i < 5; i++ {
		fmt.Println(grn, ": Hello")
	}
}
