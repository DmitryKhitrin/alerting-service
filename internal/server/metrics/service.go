package metrics

type Service interface {
	StoreMetric(metric string, name string, value string) *Error
	GetMetric(metric string, name string) (interface{}, *Error)
	GetAll() (*map[string]float64, *map[string]int64)
}
