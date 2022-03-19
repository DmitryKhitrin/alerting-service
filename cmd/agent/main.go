package main

import (
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/agent/metrics"
	"github.com/DmitryKhitrin/alerting-service/internal/agent/scheduller"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

const (
	serverPath  = "http://localhost:8080"
	contentType = "text/plain"
)

func sendStat(statString string) {
	resp, err := http.Post(serverPath+statString, contentType, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = resp.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request error with %s, http status %d", statString, resp.StatusCode)
	}
}

func collectStats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
}

func sendStats() {

	storedMetrics := metrics.CollectGauge()
	for _, metric := range storedMetrics {
		sendStat(metric.GetValue())
	}

	storedLocalMetrics := metrics.CollectCounter()
	for _, localMetric := range storedLocalMetrics {
		sendStat(localMetric.GetValue())
	}
}

func RunCollectStats() {
	scheduller.Schedule(collectStats, pollInterval)
}

func RunSendStats() {
	scheduller.Schedule(sendStats, reportInterval)
}

func registerSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
}

func main() {
	registerSignals()
	go RunCollectStats()
	RunSendStats()
}
