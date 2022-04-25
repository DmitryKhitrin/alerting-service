package metrics

import (
	"net/http"
)

type Service interface {
	StoreMetric(metric string, name string, value string) *Error
	GetMetric(metric string, name string) (interface{}, *Error)
	GetTemplateWriter() (func(w http.ResponseWriter) error, error)
}
