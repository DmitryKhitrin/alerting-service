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

	Frees := metrics.Frees(&statData.memStats)
	sendStat(Frees)

	GCSys := metrics.GCSys(&statData.memStats)
	sendStat(GCSys)

	GCCPUFraction := metrics.GCSys(&statData.memStats)
	sendStat(GCCPUFraction)

	HeapAlloc := metrics.HeapAlloc(&statData.memStats)
	sendStat(HeapAlloc)

	HeapIdle := metrics.HeapIdle(&statData.memStats)
	sendStat(HeapIdle)

	HeapInuse := metrics.HeapInuse(&statData.memStats)
	sendStat(HeapInuse)

	HeapObjects := metrics.HeapObjects(&statData.memStats)
	sendStat(HeapObjects)

	HeapReleased := metrics.HeapReleased(&statData.memStats)
	sendStat(HeapReleased)

	HeapSys := metrics.HeapSys(&statData.memStats)
	sendStat(HeapSys)

	LastGC := metrics.LastGC(&statData.memStats)
	sendStat(LastGC)

	Lookups := metrics.Lookups(&statData.memStats)
	sendStat(Lookups)

	MSpanSys := metrics.MSpanSys(&statData.memStats)
	sendStat(MSpanSys)

	MSpanInuse := metrics.MSpanInuse(&statData.memStats)
	sendStat(MSpanInuse)

	Malloc := metrics.Mallocs(&statData.memStats)
	sendStat(Malloc)

	MCacheSys := metrics.MCacheSys(&statData.memStats)
	sendStat(MCacheSys)

	MCacheInuse := metrics.MCacheInuse(&statData.memStats)
	sendStat(MCacheInuse)

	NextGC := metrics.NextGC(&statData.memStats)
	sendStat(NextGC)

	NumForcedGC := metrics.NumForcedGC(&statData.memStats)
	sendStat(NumForcedGC)

	NumGC := metrics.NumGC(&statData.memStats)
	sendStat(NumGC)

	OtherSys := metrics.OtherSys(&statData.memStats)
	sendStat(OtherSys)

	PauseTotalNs := metrics.PauseTotalNs(&statData.memStats)
	sendStat(PauseTotalNs)

	StackInuse := metrics.StackInuse(&statData.memStats)
	sendStat(StackInuse)

	StackSys := metrics.StackSys(&statData.memStats)
	sendStat(StackSys)

	Sys := metrics.Sys(&statData.memStats)
	sendStat(Sys)

	TotalAlloc := metrics.TotalAlloc(&statData.memStats)
	sendStat(TotalAlloc)

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
