package server

import (
	"context"
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/server/config"
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	metricsHandler "github.com/DmitryKhitrin/alerting-service/internal/server/metrics/handler"
	metricsService "github.com/DmitryKhitrin/alerting-service/internal/server/metrics/service"
	"github.com/DmitryKhitrin/alerting-service/internal/server/repositories"
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os/signal"
	"syscall"
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

func getRouter(a *App) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	metricsHandler.RegisterHTTPEndpoints(router, a.metricsService)
	return router
}

func LaunchServer() error {
	cfg := config.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}

	app := NewApp(&cfg)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: getRouter(app),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	go func() {
		<-ctx.Done()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	return nil
}
