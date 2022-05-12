package handler

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	"github.com/go-chi/chi"
)

func RegisterHTTPEndpoints(router *chi.Mux, s metrics.Service) {
	h := NewHandler(s)

	router.Get("/", h.GetAllHandler)

	router.Route("/value", func(r chi.Router) {
		r.Post("/", h.GetJSON)
		r.Get("/{type}/{name}", h.GetPlain)
	})

	router.Route("/update", func(r chi.Router) {
		r.Post("/", h.UpdateJSON)
		r.Post("/{type}/{name}/{value}", h.UpdatePlain)
	})

}
