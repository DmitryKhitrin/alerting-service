package repositories

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/service"
	"sync"
)

type HashMetricsRepository struct {
	gauge   map[string]float64
	counter map[string]int64
}

var hashStorage = HashMetricsRepository{
	gauge:   make(map[string]float64),
	counter: make(map[string]int64),
}
var storageMutex = &sync.RWMutex{}

func (s *HashMetricsRepository) SetGauge(name string, value float64) {
	storageMutex.Lock()
	hashStorage.gauge[name] = value
	storageMutex.Unlock()
}

func (s *HashMetricsRepository) SetCounter(name string, value int64) {
	storageMutex.Lock()
	if val, ok := hashStorage.counter[name]; ok {
		hashStorage.counter[name] = val + value
	} else {
		hashStorage.counter[name] = value
	}
	storageMutex.Unlock()
}

func GetHashStorageRepository() service.MetricsRepository {
	return &hashStorage
}
