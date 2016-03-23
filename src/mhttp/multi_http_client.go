package mhttp

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"pool"
)

var (
	httpConfig  *MHttpConfig
	hclientPool *pool.Pool
)

type MHttpConfig struct {
	RequestTimeout      int // Request timeout in millisecond
	MaxHttpClients      int // Maximum number of http client can be created in pool
	MaxIdelConnsPerHost int // Maximum idel conns per host (Keeps persitent conns)
}

type HttpRequest struct {
	req  *http.Request
	resp *http.Response
	err  error
}

type MHttpContext struct {
	reqs      []*HttpRequest
	waitGroup sync.WaitGroup
}

// Return new http.Client
func newHttpClient() (interface{}, error) {
	transport := &http.Transport{DisableKeepAlives: false,
		MaxIdleConnsPerHost: httpConfig.MaxIdelConnsPerHost,
	}

	timeout := time.Duration(time.Duration(httpConfig.RequestTimeout) * time.Millisecond)

	hclient := &http.Client{Transport: transport,
		Timeout: timeout,
	}

	return hclient, nil
}

// Init HTTP Client configuration
// reqTimeout in millisecond
func Init(reqTimeout, maxHttpClient, idelConnPerHost int) error {
	httpConfig = &MHttpConfig{
		RequestTimeout:      reqTimeout,
		MaxHttpClients:      maxHttpClient,
		MaxIdelConnsPerHost: idelConnPerHost,
	}

	hclientPool = &pool.Pool{Size: httpConfig.MaxHttpClients,
		NewObject: newHttpClient,
	}

	err := hclientPool.Init()

	return err
}

// Get new MHttpContext
func NewMHttpContext() *MHttpContext {
	requests := make([]*HttpRequest, 0)
	mHContx := &MHttpContext{reqs: requests}
	return mHContx
}

func (cntx *MHttpContext) GetRequests() []*HttpRequest {
	return cntx.reqs
}

// Add HttpRequest in MHttpContext
func (cntx *MHttpContext) AddHttpRequest(req *HttpRequest) {
	cntx.reqs = append(cntx.reqs, req)
}

// Execute all http request
func (cntx *MHttpContext) Execute() {
	for _, hreq := range cntx.reqs {
		go hreq.execute(&cntx.waitGroup)
		cntx.waitGroup.Add(1)
	}

	cntx.waitGroup.Wait()
}

// Create new http request
func NewHttpRequest(url, data, method string) (*HttpRequest, error) {
	if "GET" == method {
		// TODO: need to check ? need to be append
		url = url + data
		data = ""
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))

	if nil != err {
		return nil, err
	}

	httpRequest := &HttpRequest{
		req:  request,
		resp: nil,
		err:  nil,
	}

	return httpRequest, nil
}

// Add Request Header
func (hreq *HttpRequest) AddHeader(key, value string) {
	hreq.req.Header.Add(key, value)
}

// Get error if any
func (hreq *HttpRequest) GetError() error {
	return hreq.err
}

// Get Response Header
func (hreq *HttpRequest) GetHeader(key string) string {
	return hreq.resp.Header.Get(key)
}

// Return Response status
func (hreq *HttpRequest) GetResponseStatus() int {
	if nil == hreq.resp {
		return http.StatusRequestTimeout
	}

	return hreq.resp.StatusCode
}

// Return Response body
func (hreq *HttpRequest) GetResponseBody() ([]byte, error) {
	if nil == hreq.resp {
		return nil, errors.New("http.Response is null")
	}

	body, err := ioutil.ReadAll(hreq.resp.Body)
	if nil != err {
		return nil, err
	}

	hreq.resp.Body.Close()

	return body, nil
}

// Execute HttpRequest
func (hreq *HttpRequest) execute(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	hclient := hclientPool.Borrow().(*http.Client)
	defer hclientPool.Release(hclient)

	resp, err := hclient.Do(hreq.req)

	if nil != err {
		hreq.err = err
	}

	// set response into HttpRequest
	hreq.resp = resp
}
