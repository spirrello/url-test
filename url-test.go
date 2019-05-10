//This app is used for basic tests against an HTTP endpoint.

package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

//httpTest sends get requests to an endpoint.
func httpTest(url *string, ch chan<- string, iteration int, httpBody *string, insecure *string) {

	//if insecure flag is true skip ssl verification
	if strings.Compare(*insecure, "true") == 0 {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	//Clock the start and finish of each request
	start := time.Now()
	resp, err := http.Get(*url)

	secs := time.Since(start).Seconds()

	if err != nil {
		print(err.Error())
	} else {
		defer resp.Body.Close()
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		httpBodyStr := *httpBody
		if strings.Compare(httpBodyStr, "true") == 0 {
			ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s, http output: %s", iteration, secs, resp.Status, bodyString)
		} else {
			ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s", iteration, secs, resp.Status)
		}

	}

}

//inputValidation checks values added.
func inputValidation(url *string, request int, httpBody string) {

	//check if url is properly formatted.
	if strings.Contains(*url, "https://") {

	} else if strings.Contains(*url, "http://") {

	} else {
		log.Println("Incorrect URL format, please use http:// or https://")
		os.Exit(1)
	}

}

func main() {

	// URL and number of parallel requests to make.
	url := flag.String("url", "https://google.com", "URL and google.com by default")
	requests := flag.Int("requests", 1, "number of concurrent requests")
	httpBody := flag.String("output", "false", "flag for printing http body or not")
	insecure := flag.String("insecure", "true", "flag for when to ignore SSL errors")

	flag.Parse()

	inputValidation(url, *requests, *httpBody)

	fmt.Println("Testing with:", *url)

	//channel
	ch := make(chan string)

	//Send the requests
	for iteration := 0; iteration < *requests; iteration++ {
		go httpTest(url, ch, iteration, httpBody, insecure)
	}

	// Loop through the results
	for i := 0; i < *requests; i++ {
		log.Println(<-ch)
	}

}
