//This app is used for basic tests against an HTTP endpoint.

package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestHttpGetRequest(t *testing.T) {

	//channel
	ch := make(chan string)

	//flag for ignoring ssl warnings
	insecure := "true"
	//run 10 tests by default
	i, _ := strconv.Atoi(GetEnv("TEST_COUNT", "10"))
	for x := 0; x < i; x++ {
		testURL := GetEnv("TEST_URL", "https://yahoo.com")
		go httpGetRequest(testURL, ch, x, "false", insecure)
	}
	for x := 0; x < i; x++ {
		result := <-ch
		if !strings.Contains(result, "200 OK") {
			t.Errorf("httpGetRequest TEST FAILED")
		} else {
			fmt.Printf("%v\n", result)
		}

	}

}
