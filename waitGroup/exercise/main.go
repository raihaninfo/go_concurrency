package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://facebook.com",
	"https://twitter.com",
	"https://github.com",
	"https://w3schools.com/",
}

var wg sync.WaitGroup

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", resp.Status)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	http.HandleFunc("/", fetchStatus)
	http.ListenAndServe(":8082", nil)
}
