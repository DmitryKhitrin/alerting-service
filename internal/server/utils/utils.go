package utils

import "strings"

const (
	NameIndex  = 3
	ValueIndex = 4
)

type MetricParams struct {
	Name  string
	Value string
}

func ParseUrl(path string) MetricParams {
	name := strings.Split(path, "/")[NameIndex]
	value := strings.Split(path, "/")[ValueIndex]

	return MetricParams{name, value}
}
