package service

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
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

func (m MetricsService) GetAll() (*map[string]float64, *map[string]int64) {
	return m.repository.GetAll()
}
