package utils

import (
	"errors"
	"fmt"
	"strings"
)

const (
	NameIndex  = 3
	ValueIndex = 4
)

type MetricParams struct {
	Name  string
	Value string
}

func ParseURL(path string) (MetricParams, error) {
	name := strings.Split(path, "/")[NameIndex]

	if len(strings.Split(path, "/")) > ValueIndex {
		value := strings.Split(path, "/")[ValueIndex]
		return MetricParams{name, value}, nil
	}

	return MetricParams{name, ""}, errors.New(fmt.Sprintf("parsing url error %s", path))
}
