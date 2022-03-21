package server

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/handlers"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func getRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(handlers.GaugePath, handlers.GaugeHandler)
	mux.HandleFunc(handlers.CounterPath, handlers.CounterHandler)
	return mux
}

func LaunchServer() {
	server := &http.Server{
		Addr:    port,
		Handler: getRouter(),
	}
	log.Fatal(server.ListenAndServe())
}
