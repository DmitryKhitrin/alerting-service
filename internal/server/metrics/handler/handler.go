package handler

import (
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"text/template"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type Handler struct {
	service metrics.Service
}

func NewHandler(service metrics.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	metric := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")
	value := chi.URLParam(r, "value")

	err := h.service.StoreMetric(metric, name, value)

	if err != nil {
		http.Error(w, err.Text, err.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetMetricHandler(w http.ResponseWriter, r *http.Request) {
	metric := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")

	value, sErr := h.service.GetMetric(metric, name)

	if sErr != nil {
		http.Error(w, sErr.Text, sErr.Status)
		return
	}

	_, err := w.Write([]byte(fmt.Sprint(value)))
	if err != nil {
		return
	}
}

func (h *Handler) GetAllHandler(w http.ResponseWriter, _ *http.Request) {
	indexPage, err := os.ReadFile("internal/server/metrics/static/index.html")
	if err != nil {
		indexPage, err = os.ReadFile("index.html")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	indexTemplate := template.Must(template.New("").Parse(string(indexPage)))
	tmp := make(map[string]interface{})
	gauge, counter := h.service.GetAll()
	tmp[Gauge] = gauge
	tmp[Counter] = counter
	err = indexTemplate.Execute(w, tmp)
	if err != nil {
		log.Println(err)
		return
	}
}
