package service

import (
	"github.com/DmitryKhitrin/alerting-service/internal/common"
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	"html/template"
	"net/http"
	"os"
)

type MetricsService struct {
	repository metrics.Repository
}

func NewMetricsService(repository metrics.Repository) *MetricsService {
	return &MetricsService{
		repository: repository,
	}
}

func (m MetricsService) StoreMetric(metric *common.Metrics) *common.Error {
	switch metric.MType {
	case common.Gauge:
		if metric.Value == nil {
			return common.NewBadRequestError("wrong gauge value")
		}
		m.repository.SetValue(metric.ID, metric.Value)
	case common.Counter:
		if metric.Delta == nil {
			return common.NewBadRequestError("wrong gauge value")
		}
		m.repository.SetValue(metric.ID, metric.Delta)
	default:
		return common.NewNotImplementedError("invalid metric type")
	}
	return nil
}

func (m MetricsService) GetMetric(metric *common.Metrics) (interface{}, *common.Error) {

	switch metric.MType {
	case common.Gauge, common.Counter:
		value, err := m.repository.GetValue(metric.MType, metric.ID)
		if err != nil {
			return nil, common.NewNotFoundError("metric not found")
		}
		switch metric.MType {
		case common.Gauge:
			val := value.(float64)
			metric.Value = &val
			return val, nil
		case common.Counter:
			val := value.(int64)
			metric.Delta = &val
			return val, nil
		}
	}
	return nil, common.NewBadRequestError("invalid metric type")
}

func (m MetricsService) GetTemplateWriter() (func(w http.ResponseWriter) error, error) {
	data := m.repository.GetAll()
	indexPage, err := os.ReadFile("internal/server/metrics/static/index.html")
	indexTemplate := template.Must(template.New("").Parse(string(indexPage)))

	if err != nil {
		_, err = os.ReadFile("index.html")
		if err != nil {
			return nil, err
		}
	}

	tmp := make(map[string]interface{})
	tmp["data"] = data

	return func(w http.ResponseWriter) error {
		err = indexTemplate.Execute(w, tmp)
		return err
	}, nil

}
