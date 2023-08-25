package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const url = "https://www.google.com"

var wg sync.WaitGroup

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	wg.Add(1)
	go performRequest(&wg)

	// Wait for the spawned goroutine to complete
	wg.Wait()

	fmt.Println("Main goroutine: All goroutines finished.")
}

func performRequest(wg *sync.WaitGroup) {
	defer wg.Done() // Mark the spawned goroutine as done when it finishes

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode)

	timeTrack(time.Now(), "Run time")
}
