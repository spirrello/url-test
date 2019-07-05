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

//httpRequest processes the actual request
// func httpRequest(url string, requestType string) (resp , err string) {

// 	switch requestType {
// 	case "GET":
// 		resp, err := http.Get(url)
// 	default:
// 		resp, err := http.Get(url)
// 	}

// 	return resp, err

// }

//httpTest sends get requestCount to an endpoint.
func httpGetRequest(url string, ch chan<- string, iteration int, httpBody string, insecure *string) {

	//if insecure flag is true skip ssl verification
	if strings.Compare(*insecure, "true") == 0 {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	//Clock the start and finish of each request
	start := time.Now()
	resp, err := http.Get(url)

	secs := time.Since(start).Seconds()

	if err != nil {
		print(err.Error())
	} else {
		defer resp.Body.Close()

		//print headers
		// for k, v := range resp.Header {
		// 	fmt.Print(k)
		// 	fmt.Print(" : ")
		// 	fmt.Println(v)
		// }
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		httpBodyStr := httpBody
		if strings.Compare(httpBodyStr, "true") == 0 {
			ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s, http output: %s", iteration, secs, resp.Status, bodyString)
		} else {
			ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s", iteration, secs, resp.Status)
		}

	}

}

func httpPostRequest(url string, ch chan<- string, iteration int, httpBody string, insecure *string) {

	//if insecure flag is true skip ssl verification
	if strings.Compare(*insecure, "true") == 0 {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	//Clock the start and finish of each request
	start := time.Now()
	resp, err := http.Post(url, "application/json", nil)

	secs := time.Since(start).Seconds()

	if err != nil {
		print(err.Error())
	} else {
		defer resp.Body.Close()

		//print headers
		// for k, v := range resp.Header {
		// 	fmt.Print(k)
		// 	fmt.Print(" : ")
		// 	fmt.Println(v)
		// }
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		httpBodyStr := httpBody
		if strings.Compare(httpBodyStr, "true") == 0 {
			ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s, http output: %s", iteration, secs, resp.Status, bodyString)
		} else {
			ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s", iteration, secs, resp.Status)
		}

	}

}

//inputValidation checks values added.
func inputValidation(url *string, request int, httpBody string, requestType string) {

	//check if url is properly formatted.
	if strings.Contains(*url, "https://") {

	} else if strings.Contains(*url, "http://") {

	} else {
		log.Println("Incorrect URL format, please use http:// or https://")
		os.Exit(1)
	}

	//If the requset type is incorrect then exit.
	switch requestType {
	case "GET":
	case "POST":
	default:
		log.Println("Incorrect request-type, GET or POST accepted.")
		os.Exit(1)

	}

}

func main() {

	// Input options
	url := flag.String("url", "https://google.com", "URL and google.com by default")
	requestCount := flag.Int("request-count", 1, "number of concurrent requestCount")
	httpBody := flag.String("output", "false", "flag for printing http body or not")
	insecure := flag.String("insecure", "true", "flag for when to ignore SSL errors")
	requestType := flag.String("request-type", "GET", "GET, POST, etc")

	flag.Parse()

	inputValidation(url, *requestCount, *httpBody, *requestType)

	fmt.Println("Testing with:", *url)

	//channel
	ch := make(chan string)

	//Send the requests
	for iteration := 0; iteration < *requestCount; iteration++ {
		if strings.Compare(*requestType, "GET") == 0 {
			go httpGetRequest(*url, ch, iteration, *httpBody, insecure)

		} else if strings.Compare(*requestType, "POST") == 0 {
			go httpPostRequest(*url, ch, iteration, *httpBody, insecure)
		}

	}

	// Loop through the results
	for i := 0; i < *requestCount; i++ {
		log.Println(<-ch)
	}

}
