package storage

import "fmt"

type Gauge struct {
	Name  string
	Value float64
}

type Counter struct {
	Name  string
	Value int64
}

type Storage struct {
	gauge   map[string]float64
	counter map[string]int64
}

var storage = Storage{
	gauge:   make(map[string]float64),
	counter: make(map[string]int64),
}

func (s *Storage) SetGauge(metric *Gauge) {
	storage.gauge[metric.Name] = metric.Value
	fmt.Println(storage)
}

func (s *Storage) SetCounter(metric *Counter) {
	if val, ok := storage.counter[metric.Name]; ok {
		storage.counter[metric.Name] = val + metric.Value
	} else {
		storage.counter[metric.Name] = metric.Value
	}
	fmt.Println(storage)
}

func GetStorage() *Storage {
	return &storage
}
