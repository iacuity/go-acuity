package stathat

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"pool"
)

const (
	SHATHAT_URL      = "http://api.stathat.com/ez" // Stathat api server url
	MAX_HTTP_CLIENT  = 50                          // http.client pool size
	MAX_HTTP_TIMEOUT = 5 * time.Second             // timeout in second
	RETRY_COUNT      = 3                           // retry count if failed to send stat
)

var (
	ezAPIKey    string     // The ezkey for your StatHat account (defaults to email address)
	hclientPool *pool.Pool // http.client pool for faster and reusablility
)

// Return new http.Client
func newHttpClient() (interface{}, error) {
	transport := &http.Transport{
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 10,
	}

	timeout := time.Duration(MAX_HTTP_TIMEOUT)

	hclient := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	return hclient, nil
}

func Init(ezKey string) error {
	ezAPIKey = ezKey
	hclientPool = &pool.Pool{Size: MAX_HTTP_CLIENT,
		NewObject: newHttpClient,
	}
	err := hclientPool.Init()

	return err
}

type StatContext struct {
	EzKey  string        `json:"ezkey"`
	Metric []*statMetric `json:"data"`
}

type statMetric struct {
	Stat      string  `json:"stat"`
	Count     int     `json:"count,omitempty"`
	Value     float64 `json:"value,omitempty"`
	TimeStamp int64   `json:"t,omitempty"`
}

func NewStatContext() *StatContext {
	return &StatContext{
		EzKey:  ezAPIKey,
		Metric: make([]*statMetric, 0),
	}
}

func (cntx *StatContext) AddValueMetric(key string, value float64) {
	sMetric := &statMetric{
		Stat:      key,
		Value:     value,
		TimeStamp: time.Now().Unix(),
	}
	cntx.Metric = append(cntx.Metric, sMetric)
}

func (cntx *StatContext) AddCountMetric(key string, count int) {
	sMetric := &statMetric{
		Stat:      key,
		Count:     count,
		TimeStamp: time.Now().Unix(),
	}
	cntx.Metric = append(cntx.Metric, sMetric)
}

func (cntx *StatContext) Push() {
	data, err := json.Marshal(cntx)

	if nil != err {
		log.Println("Error: ", err.Error())
	}

	retryCount := RETRY_COUNT
	if 0 == RETRY_COUNT {
		retryCount = 1
	}

	// retry if failed to send stats
	for retryCount > 0 {
		retryCount--

		request, err := http.NewRequest("POST", SHATHAT_URL, bytes.NewBuffer(data))
		request.Header.Add("Content-Type", "application/json")

		if nil != err {
			log.Println("Error: ", err.Error())
		}

		hclient := hclientPool.Borrow().(*http.Client)
		defer hclientPool.Release(hclient)

		resp, err := hclient.Do(request)

		if nil != err {
			log.Println("Error: ", err.Error())
			continue
		}

		if nil != resp && resp.StatusCode == http.StatusOK { // ON SUCCESS, break the look
			break
		}
	}
}
