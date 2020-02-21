//This app is used for basic tests against an HTTP endpoint.

package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var url string
var requestCount int
var output string
var insecure string
var requestType string
var postFile string
var sleepTime int

func init() {

	flag.StringVar(&url, "url", "https://google.com", "URL and google.com by default")
	flag.IntVar(&requestCount, "request-count", 1, "number of concurrent requestCount")
	flag.StringVar(&output, "output", "false", "flag for printing http body or not")
	flag.StringVar(&insecure, "insecure", "true", "flag for when to ignore SSL errors")
	flag.StringVar(&requestType, "request-type", "GET", "GET, POST, etc")
	flag.StringVar(&postFile, "post-file", "", "file to post")
	flag.IntVar(&sleepTime, "sleep", 15, "Time to sleep between iterations")
}

func main() {

	flag.Parse()

	inputValidation(url, requestCount, output, requestType, postFile)

	timeStamp := getTimeStamp()
	fmt.Println(`{"level": "INFO", "timeStamp": "` + timeStamp + `", "messageType": "url", "url": "` + url + `"}`)

	ch := make(chan string)

	//Wait for SIGTERM
	setupCloseHandler()

	//Send the requests
	for {
		//Loop through the request count
		for loopCount := 0; loopCount < requestCount; loopCount++ {
			if strings.Compare(requestType, "GET") == 0 {
				go httpGetRequest(url, ch, loopCount, output, insecure)

			}
			if strings.Compare(requestType, "POST") == 0 {
				if strings.Compare(postFile, "") != 0 {
					go httpPostFileRequest(url, postFile, ch, loopCount, insecure)
				} else {
					go httpPostRequest(url, ch, loopCount, output, insecure)
				}

			}

		}

		// Loop through the results
		for i := 0; i < requestCount; i++ {
			fmt.Printf("%v\n", <-ch)
		}
		//Sleep for a duration between iterations
		if sleepTime > 0 {
			timeStamp := getTimeStamp()
			log.Println("\r" + `{"level": "INFO", "timeStamp": "` + timeStamp + `", messageType": "sleep", "duration": "` + strconv.Itoa(sleepTime) + `"}`)
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
	}

}

//getTimeStamp returns a usable time stamp
func getTimeStamp() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05.999")
}

//getFilePath fetches the absolute path to a file
func getFilePath(postFile string) string {

	//get real path to file to mitigate G304 / CWE-22
	filedirectory := filepath.Dir(postFile)
	dirPath, err := filepath.Abs(filedirectory)
	filePath := dirPath + "/" + postFile
	if err != nil {
		log.Fatal(err)
	}
	return filePath
}

//setupCloseHandler to close the program gracefully.
func setupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		timeStamp := getTimeStamp()
		log.Println("\r" + `{"level": "INFO", "timeStamp": "` + timeStamp + `", messageType": "exit"}`)
		os.Exit(0)
	}()
}

//httpTest sends get requestCount to an endpoint.
func httpGetRequest(url string, ch chan<- string, iteration int, httpBody string, insecure string) {

	//if insecure flag is true skip ssl verification
	if strings.Compare(insecure, "true") == 0 {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	//Clock the start and finish of each request
	start := time.Now()
	/* #nosec G107 */
	resp, err := http.Get(url)
	//get time stamp of when work completed
	timeStamp := getTimeStamp()

	secs := time.Since(start).Seconds()

	if err != nil {
		print(err.Error())
	} else {
		defer resp.Body.Close()

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)

		httpBodyStr := httpBody
		if strings.Compare(httpBodyStr, "true") == 0 {
			ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s, http output: %s", iteration, secs, resp.Status, bodyString)
		} else {
			//convert secs to type string and then prep result in json format
			secsStr := fmt.Sprintf("%f", secs)
			result := `{"level": "INFO", "timeStamp": "` + timeStamp + `", "messageType": "test", "test": "` + strconv.Itoa(iteration) +
				`", "totalTime": "` + secsStr + `", "statusCode": "` + resp.Status + `"}`
			ch <- result
		}

	}

}

func httpPostRequest(url string, ch chan<- string, iteration int, httpBody string, insecure string) {

	//if insecure flag is true skip ssl verification
	if strings.Compare(insecure, "true") == 0 {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	//Clock the start and finish of each request
	start := time.Now()
	/* #nosec G107 */
	resp, err := http.Post(url, "application/json", nil)

	secs := time.Since(start).Seconds()

	if err != nil {
		print(err.Error())
	} else {
		defer resp.Body.Close()

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

//httpPostRequest posts a file
func httpPostFileRequest(url string, postFile string, ch chan<- string, iteration int, insecure string) {

	//if insecure flag is true skip ssl verification
	if strings.Compare(insecure, "true") == 0 {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	//Clock the start and finish of each request
	start := time.Now()
	secs := time.Since(start).Seconds()

	//get real path to file to avoid G304
	filePath := getFilePath(postFile)
	/* #nosec G304 */
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	/* #nosec G107 */
	resp, err := http.Post(url, "application/json", file)

	if err != nil {
		log.Fatal(err)
	}

	resp.Close = true

	defer resp.Body.Close()

	ch <- fmt.Sprintf("test: %d, time spent: %.2f seconds, result: %s", iteration, secs, resp.Status)

	if err != nil {
		log.Fatal(err)
	}
}

//isJSON check for valid JSON
func isJSON(postFile string) bool {

	filePath := getFilePath(postFile)
	/* #nosec G304 */
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var js json.RawMessage
	return json.Unmarshal(byteValue, &js) == nil
}

//inputValidation checks values added.
func inputValidation(url string, request int, httpBody string, requestType string, postFile string) {

	//check if url is properly formatted.
	if strings.Contains(url, "https://") {

	} else if strings.Contains(url, "http://") {

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

	//Check for valid json
	if strings.Contains(postFile, ".json") {
		if isJSON(postFile) {

			log.Println("json payload incorrectly formatted")
			os.Exit(1)
		}
	}

}
