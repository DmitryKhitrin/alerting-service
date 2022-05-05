package server

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/config"
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	metricsHandler "github.com/DmitryKhitrin/alerting-service/internal/server/metrics/handler"
	metricsService "github.com/DmitryKhitrin/alerting-service/internal/server/metrics/service"
	"github.com/DmitryKhitrin/alerting-service/internal/server/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type App struct {
	metricsService metrics.Service
}

func NewApp(cfg *config.Config) *App {
	repository := repositories.NewLocalStorageRepository(cfg)

	return &App{
		metricsService: metricsService.NewMetricsService(repository),
	}
}

func GetRouter(a *App) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	metricsHandler.RegisterHTTPEndpoints(router, a.metricsService)
	return router
}
