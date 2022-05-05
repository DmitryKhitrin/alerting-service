package main

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server"
	"log"
)

func main() {
	if err := server.LaunchServer(); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
