package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

//HTTPMetricsRetrieval retrieve metrics from url
type HTTPMetricsRetrieval struct {
	client *http.Client
	url    string
}

//NewHTTPMetricsRetrieval HTTPMetricsRetrieval constructor
func NewHTTPMetricsRetrieval(url string, timeout int) *HTTPMetricsRetrieval {
	return &HTTPMetricsRetrieval{client: &http.Client{Timeout: time.Millisecond * time.Duration(timeout)}, url: url}
}

//Get retrieve metrics from url
func (m *HTTPMetricsRetrieval) Get() ([]byte, error) {
	resp, err := m.client.Get(m.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
