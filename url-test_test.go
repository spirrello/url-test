//This app is used for basic tests against an HTTP endpoint.

package main

import (
	"strings"
	"testing"
)

func TestHttpGetRequest(t *testing.T) {

	//channel
	ch := make(chan string)

	//flag for ignoring ssl warnings
	insecure := "true"
	go httpGetRequest("https://yahoo.com", ch, 1, "false", insecure)

	//for i := 0; i < 1; i++ {
	result := <-ch
	if !strings.Contains(result, "200 OK") {
		t.Errorf("httpGetRequest TEST FAILED")
	}

	//}

}
