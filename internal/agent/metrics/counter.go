package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/common"
	"math/rand"
)

type LocalMetrics struct {
	PollCount   int64
	RandomValue int64
}

type Counter struct {
	name  string
	value int64
}

func (g Counter) GetValue() *common.Metrics {
	return &common.Metrics{
		ID:    g.name,
		MType: common.Counter,
		Delta: &g.value,
	}
}

var localMetrics LocalMetrics

func GetCounterMetrics() []Counter {
	localMetrics.PollCount++
	localMetrics.RandomValue = rand.Int63()

	counter := []Counter{
		{name: "PollCount", value: localMetrics.PollCount},
		{name: "RandomValue", value: localMetrics.RandomValue},
	}
	return counter
}
