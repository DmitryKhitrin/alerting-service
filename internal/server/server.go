package server

import (
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/server/handlers"
	"github.com/DmitryKhitrin/alerting-service/internal/server/repositories"
	"github.com/DmitryKhitrin/alerting-service/internal/server/service"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func getRouter() *http.ServeMux {

	repository := repositories.GetHashStorageRepository()

	mux := http.NewServeMux()
	mux.HandleFunc("/update/gauge/", func(writer http.ResponseWriter, request *http.Request) {
		service.PostMetricHandler(writer, request, service.Gauge, repository)
	})
	mux.HandleFunc("/update/counter/", func(writer http.ResponseWriter, request *http.Request) {
		service.PostMetricHandler(writer, request, service.Counter, repository)
	})
	mux.HandleFunc("/update/counter", handlers.NotImplemented)
	mux.HandleFunc("/update/gauge", handlers.NotImplemented)
	mux.HandleFunc("/update/", handlers.NotImplemented)

	return mux
}

func LaunchServer() {
	server := &http.Server{
		Addr:    port,
		Handler: getRouter(),
	}
	fmt.Println("Starting on port:", port)
	log.Fatal(server.ListenAndServe())
}
