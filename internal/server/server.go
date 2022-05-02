package server

import (
	"context"
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	metricsHandler "github.com/DmitryKhitrin/alerting-service/internal/server/metrics/handler"
	metricsService "github.com/DmitryKhitrin/alerting-service/internal/server/metrics/service"
	"github.com/DmitryKhitrin/alerting-service/internal/server/repositories"
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type config struct {
	Address string `env:"ADDRESS" envDefault:"localhost:8080"`
}

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

func LaunchServer() error {
	app := NewApp()

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Println(cfg.Address)

	srv := &http.Server{
		Addr:    cfg.Address,
		Handler: getRouter(app),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return srv.Shutdown(ctx)
}
