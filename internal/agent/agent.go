package agent

import (
	"github.com/DmitryKhitrin/alerting-service/internal/agent/config"
	"github.com/DmitryKhitrin/alerting-service/internal/agent/scheduller"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunCollectStats(duration time.Duration) {
	scheduller.Schedule(CollectStats, duration)
}

func RunSendStats(address string, duration time.Duration) {
	requestService := NewRequestService(address)
	statSender := NewStatsSender(requestService)
	scheduller.Schedule(statSender.Send, duration)
}

func LaunchAgent() {
	cfg := config.NewAgentConfig()

	go RunCollectStats(cfg.PollInterval)
	go RunSendStats(cfg.Address, cfg.ReportInterval)

	cancelSignal := make(chan os.Signal, 1)
	signal.Notify(cancelSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-cancelSignal
	os.Exit(1)
}
