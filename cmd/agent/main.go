package main

import (
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/agent/metrics/alloc"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 2 * time.Second
)

const (
	serverPath  = "http://localhost:8080"
	contentType = "text/plain"
)

type StatData struct {
	mu          sync.Mutex
	memStats    runtime.MemStats
	PollCount   int64
	RandomValue int
}

var statData StatData

func sendStat(statString string) {
	resp, err := http.Post(serverPath+statString, contentType, nil)
	fmt.Println(statString)
	//defer resp.Body.Close()

	fmt.Println(resp, err)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request error with %s, http status %d", statString, resp.StatusCode)
	}
}

func collectStats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	statData.mu.Lock()
	statData.PollCount++
	statData.RandomValue = 12
	statData.memStats = memStats
	statData.mu.Unlock()
}

func sendStats() {
	Alloc := alloc.Get(statData.memStats)
	sendStat(Alloc)
}

func RunCollectStats() {
	ticker := time.NewTicker(pollInterval)

	for {
		<-ticker.C
		collectStats()
	}
}

func RunSendStats() {
	ticker := time.NewTicker(reportInterval)

	for {
		<-ticker.C
		sendStats()
	}
}

func main() {
	go RunCollectStats()
	RunSendStats()
}
