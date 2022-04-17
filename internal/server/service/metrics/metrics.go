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
	SetValue(name string, value interface{})
	GetValue(metric string, name string) (interface{}, error)
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
		ctx.Storage.SetValue(name, value)
	case Counter:
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			http.Error(w, "wrong counter value", http.StatusBadRequest)
			return
		}
		ctx.Storage.SetValue(name, value)
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
	case Gauge, Counter:
		value, err := ctx.Storage.GetValue(metric, name)
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
