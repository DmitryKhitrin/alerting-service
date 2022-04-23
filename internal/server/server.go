package server

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/metricks"
	"github.com/DmitryKhitrin/alerting-service/internal/server/repositories"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
)

const (
	defaultPort = "8080"
)

func getRouter() *chi.Mux {

	handler := metricks.NewHandler(repositories.NewHashStorageRepository())

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", handler.GetAllHandler)

	router.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", handler.PostHandler)
	})
	router.Route("/value", func(r chi.Router) {
		r.Get("/{type}/{name}", handler.GetMetricHandler)
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
