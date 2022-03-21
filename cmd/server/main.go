package main

import (
	"github.com/DmitryKhitrin/alerting-service/internal/common"
	"github.com/DmitryKhitrin/alerting-service/internal/server"
)

func main() {
	server.LaunchServer()
	common.RegisterCancelSignals()
}
