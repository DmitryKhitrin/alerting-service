package server

import (
	"context"
	"github.com/DmitryKhitrin/alerting-service/internal/server/repositories"
	"github.com/DmitryKhitrin/alerting-service/internal/server/service/metrics"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func MetricCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), MetricCtx, &metrics.Ctx{
			Storage: repositories.GetHashStorageRepository(),
		})
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func getRouter() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/update", func(r chi.Router) {
		r.Use(MetricCtx)
		r.Post("/{type}/{name}/{value}", metrics.PostMetricHandler)
	})
	return router
}

func LaunchServer() {
	log.Fatal(http.ListenAndServe(port, getRouter()))
}
