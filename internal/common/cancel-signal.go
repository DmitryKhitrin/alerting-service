package common

import (
	"os"
	"os/signal"
	"syscall"
)

func RegisterCancelSignals() {
	cancelSignal := make(chan os.Signal, 1)
	signal.Notify(cancelSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-cancelSignal
	os.Exit(1)
}
