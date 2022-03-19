package main

import (
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/agent/metrics"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
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
	RandomValue int64
}

var statData StatData

func sendStat(statString string) {
	resp, err := http.Post(serverPath+statString, contentType, nil)

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
	statData.RandomValue = rand.Int63()
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

	PullCount := metrics.PullCount(statData.PollCount)
	sendStat(PullCount)

	RandomValue := metrics.RandomValue(statData.RandomValue)
	sendStat(RandomValue)
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

func registerSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
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
