package server

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	metricsHandler "github.com/DmitryKhitrin/alerting-service/internal/server/metrics/handler"
	metricsService "github.com/DmitryKhitrin/alerting-service/internal/server/metrics/service"
	"github.com/DmitryKhitrin/alerting-service/internal/server/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

const (
	defaultPort = "8080"
)

type App struct {
	metricsService metrics.Service
}

func NewApp() *App {
	repository := repositories.NewLocalStorageRepository()

	return &App{
		metricsService: metricsService.NewMetricsService(repository),
	}
}

func getRouter(a *App) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	metricsHandler.RegisterHTTPEndpoints(router, a.metricsService)
	return router
}

func LaunchServer() {
	app := NewApp()

	svr := &http.Server{
		Addr:    "localhost" + ":" + defaultPort,
		Handler: getRouter(app),
	}
	svr.SetKeepAlivesEnabled(false)
	log.Printf("listening on port " + defaultPort)
	log.Fatal(svr.ListenAndServe())
}
