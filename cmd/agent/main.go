package main

import (
	"github.com/DmitryKhitrin/alerting-service/internal/agent"
	"github.com/DmitryKhitrin/alerting-service/internal/agent/scheduller"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func RunCollectStats() {
	scheduller.Schedule(agent.CollectStats, pollInterval)
}

func RunSendStats() {
	scheduller.Schedule(agent.SendStats, reportInterval)
}

func main() {
	go RunCollectStats()
	go RunSendStats()

	cancelSignal := make(chan os.Signal, 1)
	signal.Notify(cancelSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-cancelSignal
	os.Exit(1)
}
