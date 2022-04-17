package repositories

import (
	"errors"
	"github.com/DmitryKhitrin/alerting-service/internal/server/service/metrics"
	"sync"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type HashRepository struct {
	gauge   map[string]float64
	counter map[string]int64
}

var hashRepository = HashRepository{
	gauge:   make(map[string]float64),
	counter: make(map[string]int64),
}
var storageMutex = &sync.RWMutex{}

func (s *HashRepository) SetGauge(name string, value float64) {
	hashRepository.gauge[name] = value
}

func (s *HashRepository) SetCounter(name string, value int64) {
	if val, ok := hashRepository.counter[name]; ok {
		hashRepository.counter[name] = val + value
	} else {
		hashRepository.counter[name] = value
	}
}

func (s *HashRepository) SetValue(name string, value interface{}) {
	storageMutex.Lock()
	switch v2 := value.(type) {
	case int64:
		s.SetCounter(name, v2)
	case float64:
		s.SetGauge(name, v2)
	default:
	}
	storageMutex.Unlock()
}

func (s *HashRepository) GetGauge(name string) (float64, error) {
	storageMutex.Lock()
	value, ok := hashRepository.gauge[name]
	storageMutex.Unlock()
	if !ok {
		return value, errors.New("invalid metric name")
	}
	return value, nil
}

func (s *HashRepository) GetCounter(name string) (int64, error) {
	storageMutex.Lock()
	value, ok := hashRepository.counter[name]
	storageMutex.Unlock()
	if !ok {
		return value, errors.New("invalid metric name")
	}
	return value, nil
}

func (s *HashRepository) GetValue(metric string, name string) (interface{}, error) {
	switch metric {
	case Gauge:
		return s.GetGauge(name)
	case Counter:
		return s.GetCounter(name)
	default:
		return "", errors.New("invalid metric type")
	}
}

func (s *HashRepository) GetAll() (*map[string]float64, *map[string]int64) {
	return &hashRepository.gauge, &hashRepository.counter
}

func GetHashStorageRepository() metrics.Repository {
	return &hashRepository
}
