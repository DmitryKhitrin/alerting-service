package utils

import "fmt"

func MakeStatStringGauge(name string, value float64) string {
	return fmt.Sprintf("/update/gauge/%s/%G", name, value)
}

func MakeStatStringCounter(name string, value int64) string {
	return fmt.Sprintf("/update/counter/%s/%d", name, value)
}
