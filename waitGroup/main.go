package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func myFunc() {
	time.Sleep(1 + time.Second)
	fmt.Println("Finished Executing Goroutine")
	wg.Done()
}

func main() {
	fmt.Println("Go wait Group")
	wg.Add(1)
	go myFunc()
	wg.Wait()
	fmt.Println("Finished Executing my go program")
}
