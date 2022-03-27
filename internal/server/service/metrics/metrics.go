package metrics

import (
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

const (
	Gauge     = "gauge"
	Counter   = "counter"
	MetricCtx = "MetricsCtx"
)

type Ctx struct {
	Storage Repository
}

type Repository interface {
	SetCounter(name string, value int64)
	SetGauge(name string, value float64)
}

func PostMetricHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	metricCtx := ctx.Value(MetricCtx).(*Ctx)
	metric := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")
	value := chi.URLParam(r, "value")

	switch metric {
	case Gauge:
		value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			http.Error(w, "wrong gauge value", http.StatusBadRequest)
			return
		}
		metricCtx.Storage.SetGauge(name, value)
	case Counter:
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			http.Error(w, "wrong counter value", http.StatusBadRequest)
			return
		}
		metricCtx.Storage.SetCounter(name, value)
	default:
		http.Error(w, "invalid metric type", http.StatusNotImplemented)
		return
	}
	w.WriteHeader(http.StatusOK)
}
