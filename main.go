package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"

	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/tidwall/gjson"
)

var logInfo = log.New(os.Stdout, "INFO: ", log.Ldate|log.Lmicroseconds)
var logError = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Lmicroseconds)

func main() {
	serviceURL := flag.String("u", "", "service url with metrics used to send to statds")
	serviceTimeout := flag.Int("t", 1000, "service timeout in millis")
	mapPath := flag.String("m", "map.yml", "map file from json to statsd key")
	interval := flag.Int("i", 60, "interval in seconds to collect metrics")
	statsdAddr := flag.String("s", "127.0.0.1:8125", "statsd address")
	statsdPrefix := flag.String("p", "", "statsd key prefix")
	flag.Parse()

	if *serviceURL == "" {
		logError.Fatalf("service url is empty, use -h for help")
	}

	templateRender := NewTemplateRender()
	template, err := templateRender.Render(readFile(*mapPath))
	if err != nil {
		logError.Fatalf("could not render statsd map template, err:[%v]", err)
	}

	statsdMap, err := getMap(template)
	if err != nil {
		logError.Fatalf("could not parse statsd map, err:[%v]", err)
	}

	metricsRetrieval := NewHTTPMetricsRetrieval(*serviceURL, *serviceTimeout)

	statsd, err := NewStatsdSender(*statsdAddr, *statsdPrefix)
	if err != nil {
		logError.Fatalf("could not create client for statsd:[%s], err:[%v]", *statsdAddr, err)
	}
	defer statsd.Close()

	gocron.Every(uint64(*interval)).Seconds().Do(execute, metricsRetrieval, statsd, statsdMap)
	<-gocron.Start()
}

func execute(metricsRetrieval *HTTPMetricsRetrieval, statsd *StatsdSender, statsdMap map[string]string) {
	start := time.Now()
	metrics, err := metricsRetrieval.Get()
	if err != nil {
		logError.Printf("could not get metrics from service, err:[%v]", err)
		return
	}
	logInfo.Printf("got metrics from service on %v", millis(time.Since(start)))

	start = time.Now()
	for key, value := range statsdMap {
		statsd.GaugeString(key, gjson.GetBytes(metrics, value).String())
	}
	logInfo.Printf("sent %d metrics to statsd on %v", len(statsdMap), millis(time.Since(start)))
}

func readFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		logError.Fatalf("could not read statsd map, err:[%v]", err)
	}
	return string(content)
}

func getMap(content string) (map[string]string, error) {
	m := make(map[string]string)
	if err := yaml.Unmarshal([]byte(content), &m); err != nil {
		return nil, err
	}
	return m, nil
}

func millis(d time.Duration) int {
	return int(d.Nanoseconds() / 1e6)
}
