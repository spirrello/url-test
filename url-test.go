package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func httpTest(url *string, ch chan<- string, iteration int, outputOpt *string) {
	//function for testing URLs
	//url for target
	//ch for channel
	//iteration used for printing the test iteration
	//output used for whether not to print http response

	//Disable security check
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
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

		outputStr := *outputOpt
		if strings.Compare(outputStr, "true") == 0 {
			ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s, http output: %s", iteration, secs, resp.Status, bodyString)
		} else {
			ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s", iteration, secs, resp.Status)
		}

	}

}

func main() {

	// URL and number of parallel requests to make.
	url := flag.String("url", "https://google.com", "URL and google.com by default")
	requests := flag.Int("requests", 1, "number of concurrent requests")
	output := flag.String("output", "false", "flag for printing output or not")

	flag.Parse()

	fmt.Println("Testing with:", *url)

	urlOpt := url
	requestOpt := requests
	outputOpt := output

	ch := make(chan string)

	//Send the requests
	for iteration := 0; iteration < *requestOpt; iteration++ {
		go httpTest(urlOpt, ch, iteration, outputOpt)

	}
	// Loop through the results
	for i := 0; i < *requestOpt; i++ {
		log.Println(<-ch)
	}

}
