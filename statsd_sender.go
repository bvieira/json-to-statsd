package main

import (
	"fmt"
	"time"

	"github.com/cactus/go-statsd-client/statsd"
)

//StatsdSender send content to statsd
type StatsdSender struct {
	statsdClient statsd.Statter
}

//NewStatsdSender StatsdSender constructor
func NewStatsdSender(addr, prefix string) (*StatsdSender, error) {
	statsdClient, err := statsd.NewBufferedClient(addr, prefix, 100*time.Millisecond, 512)
	if err != nil {
		return nil, err
	}
	return &StatsdSender{statsdClient: statsdClient}, nil
}

//GaugeString send string
func (s *StatsdSender) GaugeString(key, value string) error {
	return s.statsdClient.Raw(key, fmt.Sprintf("%s|g", value), 1.0)
}

//GaugeFloat send float
func (s *StatsdSender) GaugeFloat(key string, value float64) error {
	return s.statsdClient.Raw(key, fmt.Sprintf("%f|g", value), 1.0)
}

//GaugeInt send int
func (s *StatsdSender) GaugeInt(key string, value int) error {
	return s.statsdClient.Raw(key, fmt.Sprintf("%d|g", value), 1.0)
}

//Close close resources
func (s *StatsdSender) Close() {
	s.statsdClient.Close()
}
