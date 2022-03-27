package repositories

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/service/metrics"
	"sync"
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
	storageMutex.Lock()
	hashRepository.gauge[name] = value
	storageMutex.Unlock()
}

func (s *HashRepository) SetCounter(name string, value int64) {
	storageMutex.Lock()
	if val, ok := hashRepository.counter[name]; ok {
		hashRepository.counter[name] = val + value
	} else {
		hashRepository.counter[name] = value
	}
	storageMutex.Unlock()
}

func GetHashStorageRepository() metrics.Repository {
	return &hashRepository
}
