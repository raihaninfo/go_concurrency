package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// real	0m4.877s
// user	0m0.645s
// sys	0m0.186s

func sendRequest(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%d] %s\n", res.StatusCode, url)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go <url1> <url2> ....")
	}
	for _, url := range os.Args[1:] {
		sendRequest("https://" + url)
	}
}
