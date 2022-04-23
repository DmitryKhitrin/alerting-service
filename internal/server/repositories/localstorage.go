package repositories

import (
	"errors"
	"sync"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type LocalStorageRepository struct {
	mutex   sync.RWMutex
	gauge   map[string]float64
	counter map[string]int64
}

var localStorageRepository = LocalStorageRepository{
	gauge:   make(map[string]float64),
	counter: make(map[string]int64),
}

func NewLocalStorageRepository() *LocalStorageRepository {
	return &localStorageRepository
}

func (s LocalStorageRepository) setGauge(name string, value float64) {
	localStorageRepository.gauge[name] = value
}

func (s LocalStorageRepository) setCounter(name string, value int64) {
	if val, ok := localStorageRepository.counter[name]; ok {
		localStorageRepository.counter[name] = val + value
	} else {
		localStorageRepository.counter[name] = value
	}
}

func (s LocalStorageRepository) SetValue(name string, value interface{}) {
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

func (s LocalStorageRepository) getGauge(name string) (float64, error) {
	s.mutex.Lock()
	value, ok := localStorageRepository.gauge[name]
	s.mutex.Unlock()
	if !ok {
		return value, errors.New("invalid metric name")
	}
	return value, nil
}

func (s LocalStorageRepository) getCounter(name string) (int64, error) {
	s.mutex.Lock()
	value, ok := localStorageRepository.counter[name]
	s.mutex.Unlock()
	if !ok {
		return value, errors.New("invalid metric name")
	}
	return value, nil
}

func (s LocalStorageRepository) GetValue(metric string, name string) (interface{}, error) {
	switch metric {
	case Gauge:
		return s.getGauge(name)
	case Counter:
		return s.getCounter(name)
	default:
		return "", errors.New("invalid metric type")
	}
}

func (s *LocalStorageRepository) GetAll() (*map[string]float64, *map[string]int64) {
	return &localStorageRepository.gauge, &localStorageRepository.counter
}
