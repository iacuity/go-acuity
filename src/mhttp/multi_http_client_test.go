package mhttp_test

import (
	"mhttp"
	"net/http"

	"testing"
)

func TestHttp(t *testing.T) {
	requestTimeout := 1500 // In millisecond
	maxHttpClient := 100
	idelConnPerHost := 20
	err := mhttp.Init(requestTimeout, maxHttpClient, idelConnPerHost)

	if nil != err {
		t.Fatal("Failed to Init mhttp")
	}

	mhttpCntx := mhttp.NewMHttpContext()

	req1, _ := mhttp.NewHttpRequest("https://www.google.co.in/", "", "GET")
	req2, _ := mhttp.NewHttpRequest("http://www.bing.com/", "", "GET")
	req3, _ := mhttp.NewHttpRequest("https://www.facebook.com/", "", "GET")

	mhttpCntx.AddHttpRequest(req1)
	mhttpCntx.AddHttpRequest(req2)
	mhttpCntx.AddHttpRequest(req3)

	mhttpCntx.Execute()

	for _, hreq := range mhttpCntx.GetRequests() {
		if http.StatusOK != hreq.GetResponseStatus() {
			t.Fatal("Invalid request status")
		}
	}
}
