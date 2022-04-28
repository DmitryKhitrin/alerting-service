package common

import (
	"strconv"
)

const (
	Gauge   = "gauge"
	Counter = "counter"
)

type Metrics struct {
	ID    string   `json:"id"`              // имя метрики
	MType string   `json:"type"`            // параметр, принимающий значение gauge или counter
	Delta *int64   `json:"delta,omitempty"` // значение метрики в случае передачи counter
	Value *float64 `json:"value,omitempty"` // значение метрики в случае передачи gauge
}

func (m *Metrics) CreateMetric(name string, mType string, value string) *Error {
	m.ID = name
	switch mType {
	case Gauge:
		value, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return NewBadRequestError("wrong gauge value")
		}
		m.MType = Gauge
		m.Value = &value
	case Counter:
		value, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return NewBadRequestError("wrong counter value")
		}
		m.MType = Counter
		m.Delta = &value
	default:
		return NewNotImplementedError("invalid metric type")
	}
	return nil

}

func (m *Metrics) FromNameAndType(name string, mType string) {
	m.ID = name
	m.MType = mType
}
