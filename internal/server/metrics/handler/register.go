package handler

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	"github.com/go-chi/chi"
)

func RegisterHTTPEndpoints(router *chi.Mux, s metrics.Service) {
	h := NewHandler(s)

	router.Get("/", h.GetAllHandler)
	router.Post("/value", h.JSONGetMetricHandler)
	router.Post("/update", h.JSONPostHandler)

	router.Route("/update", func(r chi.Router) {
		r.Post("/{type}/{name}/{value}", h.PostHandler)
	})

	router.Route("/value", func(r chi.Router) {
		r.Get("/{type}/{name}", h.GetMetricHandler)
	})

}
