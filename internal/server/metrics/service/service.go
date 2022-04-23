package service

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type MetricsService struct {
	repository metrics.Repository
}

func NewMetricsService(repository metrics.Repository) *MetricsService {
	return &MetricsService{
		repository: repository,
	}
}

func (m MetricsService) StoreMetric(metric string, name string, value string) *metrics.Error {
	switch metric {
	case Gauge:
		value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return metrics.NewBadRequestError("wrong gauge value")
		}
		m.repository.SetValue(name, value)
	case Counter:
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return metrics.NewBadRequestError("wrong counter value")
		}
		m.repository.SetValue(name, value)
	default:
		return metrics.NewNotImplementedError("invalid metric type")
	}
	return nil
}

func (m MetricsService) GetMetric(metric string, name string) (interface{}, *metrics.Error) {
	switch metric {
	case Gauge, Counter:
		value, err := m.repository.GetValue(metric, name)
		if err != nil {
			return nil, metrics.NewNotFoundError("")
		}
		return value, nil
	default:
		return nil, metrics.NewBadRequestError("invalid metric type")
	}
}

func (m MetricsService) GetTemplateWriter() (func(w http.ResponseWriter) error, error) {
	gauge, counter := m.repository.GetAll()
	indexPage, err := os.ReadFile("internal/server/metrics/static/index.html")
	indexTemplate := template.Must(template.New("").Parse(string(indexPage)))

	if err != nil {
		_, err = os.ReadFile("index.html")
		if err != nil {
			return nil, err
		}
	}

	tmp := make(map[string]interface{})
	tmp[Gauge] = gauge
	tmp[Counter] = counter

	return func(w http.ResponseWriter) error {
		err = indexTemplate.Execute(w, tmp)
		return err
	}, nil

}
