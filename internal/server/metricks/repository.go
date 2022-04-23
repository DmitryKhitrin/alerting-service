package metricks

type MetricksRepository interface {
	SetValue(name string, value interface{})
	GetValue(metric string, name string) (interface{}, error)
	GetAll() (*map[string]float64, *map[string]int64)
}
