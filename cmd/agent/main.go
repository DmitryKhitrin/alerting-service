package main

import (
	"github.com/DmitryKhitrin/alerting-service/internal/agent"
	"github.com/DmitryKhitrin/alerting-service/internal/agent/scheduller"
	"github.com/DmitryKhitrin/alerting-service/internal/common"
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
	common.RegisterCancelSignals()
}
