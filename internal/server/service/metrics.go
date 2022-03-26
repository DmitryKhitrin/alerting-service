package service

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/utils"
	"log"
	"net/http"
	"strconv"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type MetricsRepository interface {
	SetCounter(name string, value int64)
	SetGauge(name string, value float64)
}

func PostMetricHandler(w http.ResponseWriter, r *http.Request, metricType string, repository MetricsRepository) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	params, err := utils.ParseURL(r.RequestURI)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var errors error = nil

	switch metricType {
	case Gauge:
		errors = ConvertAndStoreGauge(params.Name, params.Value, repository)
	case Counter:
		errors = ConvertAndStoreCounter(params.Name, params.Value, repository)
	}

	if errors != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func ConvertAndStoreCounter(name string, value string, repository MetricsRepository) error {
	if value, err := strconv.ParseInt(value, 10, 64); err == nil {
		repository.SetCounter(name, value)
		return nil
	} else {
		return err
	}
}

func ConvertAndStoreGauge(name string, value string, repository MetricsRepository) error {
	if value, err := strconv.ParseFloat(value, 64); err == nil {
		repository.SetGauge(name, value)
		return nil
	} else {
		return err
	}
}
