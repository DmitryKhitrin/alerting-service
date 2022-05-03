package metrics

type Repository interface {
	SetValue(name string, value interface{})
	GetValue(metric string, name string) (interface{}, error)
	GetAll() *map[string]float64
}
