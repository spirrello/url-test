package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func httpTest(url *string, ch chan<- string, i int) {
	//function for testing URLs

	//Clock the start and finish of each request
	start := time.Now()
	resp, err := http.Get(*url)
	secs := time.Since(start).Seconds()

	if err != nil {
		print(err.Error())
	} else {
		ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s", i, secs, resp.Status)
	}

}

func main() {

	// URL and number of parallel requests to make.
	url := flag.String("url", "https://google.com", "URL and google.com by default")
	requests := flag.Int("requests", 1, "number of concurrent requests")

	flag.Parse()

	fmt.Println("Testing with:", *url)

	urlOpt := url
	requestOpt := requests

	ch := make(chan string)

	//Send the requests
	for i := 0; i < *requestOpt; i++ {
		go httpTest(urlOpt, ch, i)

	}
	// Loop through the results
	for i := 0; i < *requestOpt; i++ {
		log.Println(<-ch)
	}

}
