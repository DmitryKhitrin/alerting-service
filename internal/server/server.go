package server

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/repositories"
	"github.com/DmitryKhitrin/alerting-service/internal/server/service/metrics"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

const (
	defaultPort = "8080"
)

func getRouter() *chi.Mux {

	ctx := &metrics.Ctx{
		Storage: repositories.GetHashStorageRepository(),
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		metrics.GetAllHandler(ctx, w, r)
	})

	router.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", func(w http.ResponseWriter, r *http.Request) {
			metrics.PostMetricHandler(ctx, w, r)
		})
	})
	router.Route("/value", func(r chi.Router) {
		r.Get("/{type}/{name}", func(w http.ResponseWriter, r *http.Request) {
			metrics.GetMetricHandler(ctx, w, r)
		})
	})
	return router
}

func LaunchServer() {
	svr := &http.Server{
		Addr:    "localhost" + ":" + defaultPort,
		Handler: getRouter(),
	}
	svr.SetKeepAlivesEnabled(false)
	log.Printf("listening on port " + defaultPort)
	log.Fatal(svr.ListenAndServe())
}
