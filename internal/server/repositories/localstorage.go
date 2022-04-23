package repositories

import (
	"errors"
	"github.com/DmitryKhitrin/alerting-service/internal/server/metrics"
	"sync"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type LocalStorageMetricRepository struct {
	mutex   sync.RWMutex
	gauge   map[string]float64
	counter map[string]int64
}

var hashRepository = LocalStorageMetricRepository{
	gauge:   make(map[string]float64),
	counter: make(map[string]int64),
}

func NewHashStorageRepository() metrics.Repository {
	return &hashRepository
}

func (s *LocalStorageMetricRepository) setGauge(name string, value float64) {
	hashRepository.gauge[name] = value
}

func (s *LocalStorageMetricRepository) setCounter(name string, value int64) {
	if val, ok := hashRepository.counter[name]; ok {
		hashRepository.counter[name] = val + value
	} else {
		hashRepository.counter[name] = value
	}
}

func (s *LocalStorageMetricRepository) SetValue(name string, value interface{}) {
	s.mutex.Lock()
	switch v2 := value.(type) {
	case int64:
		s.setCounter(name, v2)
	case float64:
		s.setGauge(name, v2)
	default:
	}
	s.mutex.Unlock()
}

func (s *LocalStorageMetricRepository) getGauge(name string) (float64, error) {
	s.mutex.Lock()
	value, ok := hashRepository.gauge[name]
	s.mutex.Unlock()
	if !ok {
		return value, errors.New("invalid metric name")
	}
	return value, nil
}

func (s *LocalStorageMetricRepository) getCounter(name string) (int64, error) {
	s.mutex.Lock()
	value, ok := hashRepository.counter[name]
	s.mutex.Unlock()
	if !ok {
		return value, errors.New("invalid metric name")
	}
	return value, nil
}

func (s *LocalStorageMetricRepository) GetValue(metric string, name string) (interface{}, error) {
	switch metric {
	case Gauge:
		return s.getGauge(name)
	case Counter:
		return s.getCounter(name)
	default:
		return "", errors.New("invalid metric type")
	}
}

func (s *LocalStorageMetricRepository) GetAll() (*map[string]float64, *map[string]int64) {
	return &hashRepository.gauge, &hashRepository.counter
}
