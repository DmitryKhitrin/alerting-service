package metrics

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type Ctx struct {
	Storage Repository
}

type Repository interface {
	SetCounter(name string, value int64)
	SetGauge(name string, value float64)
	GetCounter(name string) (int64, error)
	GetGauge(name string) (float64, error)
	GetAll() (*map[string]float64, *map[string]int64)
}

func PostMetricHandler(ctx *Ctx, w http.ResponseWriter, r *http.Request) {
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
		ctx.Storage.SetGauge(name, value)
	case Counter:
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			http.Error(w, "wrong counter value", http.StatusBadRequest)
			return
		}
		ctx.Storage.SetCounter(name, value)
	default:
		http.Error(w, "invalid metric type", http.StatusNotImplemented)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetMetricHandler(ctx *Ctx, w http.ResponseWriter, r *http.Request) {
	metric := chi.URLParam(r, "type")
	name := chi.URLParam(r, "name")

	switch metric {
	case Gauge:
		value, err := ctx.Storage.GetGauge(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		_, err = w.Write([]byte(fmt.Sprint(value)))
		if err != nil {
			return
		}
	case Counter:
		value, err := ctx.Storage.GetCounter(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		_, err = w.Write([]byte(fmt.Sprint(value)))
		if err != nil {
			return
		}
	default:
		http.Error(w, "invalid metric type", http.StatusBadRequest)
		return
	}
}

func GetAllHandler(ctx *Ctx, w http.ResponseWriter, _ *http.Request) {
	indexPage, err := os.ReadFile("internal/server/service/metrics/static/index.html")
	if err != nil {
		indexPage, err = os.ReadFile("index.html")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	indexTemplate := template.Must(template.New("").Parse(string(indexPage)))
	tmp := make(map[string]interface{})
	gauge, counter := ctx.Storage.GetAll()
	tmp[Gauge] = gauge
	tmp[Counter] = counter
	err = indexTemplate.Execute(w, tmp)
	if err != nil {
		log.Println(err)
		return
	}
}
