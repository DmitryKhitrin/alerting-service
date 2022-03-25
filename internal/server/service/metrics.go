package service

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/utils"
	"log"
	"net/http"
	"strconv"
)

type MetricsRepository interface {
	SetCounter(name string, value int64)
	SetGauge(name string, value float64)
}

func CounterHandler(w http.ResponseWriter, r *http.Request, storage MetricsRepository) {
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

	if value, err := strconv.ParseInt(params.Value, 10, 64); err == nil {
		storage.SetCounter(params.Name, value)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}

func GaugeHandler(w http.ResponseWriter, r *http.Request, storage MetricsRepository) {
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

	if value, err := strconv.ParseFloat(params.Value, 64); err == nil {
		storage.SetGauge(params.Name, value)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
