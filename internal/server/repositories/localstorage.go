package repositories

import (
	"errors"
	"fmt"
	"github.com/DmitryKhitrin/alerting-service/internal/common"
	"github.com/DmitryKhitrin/alerting-service/internal/server/config"
	"sync"
)

type LocalStorageRepository struct {
	mutex      *sync.RWMutex
	repository map[string]float64
	cfg        *config.Config
}

func NewLocalStorageRepository(cfg *config.Config) *LocalStorageRepository {

	repository := &LocalStorageRepository{
		mutex:      &sync.RWMutex{},
		repository: make(map[string]float64),
		cfg:        cfg,
	}
	repository.TryRestore()
	repository.RunDataDumper()
	return repository
}

func (l *LocalStorageRepository) setGauge(name string, value float64) {
	l.repository[name] = value
}

func (l *LocalStorageRepository) setCounter(name string, value int64) {

	if val, ok := l.repository[name]; ok {
		l.repository[name] = val + float64(value)
	} else {
		l.repository[name] = float64(value)
	}
}

func (l *LocalStorageRepository) SetValue(name string, value interface{}) {
	l.mutex.Lock()
	switch v2 := value.(type) {
	case *int64:
		l.setCounter(name, *v2)
	case *float64:
		l.setGauge(name, *v2)
	default:
	}
	fmt.Println()
	l.mutex.Unlock()

	if l.cfg.StoreInterval.Seconds() == 0 {
		l.SaveToFile()
	}
}

func (l *LocalStorageRepository) getGauge(name string) (float64, error) {
	value, ok := l.repository[name]
	if !ok {
		return value, errors.New("invalid metric name")
	}
	return value, nil
}

func (l *LocalStorageRepository) getCounter(name string) (int64, error) {
	v, ok := l.repository[name]
	if !ok {
		return 0, errors.New("invalid metric name")
	}
	value := int64(v)
	return value, nil
}

func (l *LocalStorageRepository) GetValue(metric string, name string) (interface{}, error) {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	switch metric {
	case common.Gauge:
		return l.getGauge(name)
	case common.Counter:
		return l.getCounter(name)
	default:
		return "", errors.New("invalid metric type")
	}
}

func (l *LocalStorageRepository) GetAll() *map[string]float64 {
	return &l.repository
}
