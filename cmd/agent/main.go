package main

import (
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/agent/metrics"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"sync"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

const (
	port        = 8080
	serverPath  = "http://localhost" + string(port)
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
	defer resp.Body.Close()

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
	statData.RandomValue = rand.Int()
	statData.PollCount++
	statData.memStats = memStats
	statData.mu.Unlock()
}

func sendStats() {
	Alloc := metrics.Alloc(&statData.memStats)
	sendStat(Alloc)

	BuckHashSys := metrics.BuckHashSys(&statData.memStats)
	sendStat(BuckHashSys)
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
