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
	msg := "hello"
	go func(m string) {
		fmt.Println(m)
	}(msg)
	msg = "world"
	fmt.Println("Main End")

	time.Sleep(100 * time.Millisecond)
}

func hello(grn int) {
	for i := 0; i < 3; i++ {
		fmt.Println(grn, ": Hello")
	}
}
