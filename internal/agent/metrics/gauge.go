package metrics

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

type Gauge struct {
	name  string
	value float64
}

func (g Gauge) GetValue() string {
	return fmt.Sprintf("/update/gauge/%s/%G", g.name, g.value)
}

func CollectGauge() []Gauge {
	var rtm runtime.MemStats
	log.Println("ticker CollectRuntimeMetrics")
	runtime.ReadMemStats(&rtm)
	rand.Seed(time.Now().Unix())
	Metrics := []Gauge{
		{name: "Alloc", value: float64(rtm.Alloc)},
		{name: "BuckHashSys", value: float64(rtm.BuckHashSys)},
		{name: "Frees", value: float64(rtm.Frees)},
		{name: "GCCPUFraction", value: rtm.GCCPUFraction},
		{name: "GCSys", value: float64(rtm.GCSys)},
		{name: "HeapAlloc", value: float64(rtm.HeapAlloc)},
		{name: "HeapIdle", value: float64(rtm.HeapIdle)},
		{name: "HeapInuse", value: float64(rtm.HeapInuse)},
		{name: "HeapObjects", value: float64(rtm.HeapObjects)},
		{name: "HeapReleased", value: float64(rtm.HeapReleased)},
		{name: "HeapSys", value: float64(rtm.HeapSys)},
		{name: "LastGC", value: float64(rtm.LastGC)},
		{name: "Lookups", value: float64(rtm.Lookups)},
		{name: "MCacheInuse", value: float64(rtm.MCacheInuse)},
		{name: "MCacheSys", value: float64(rtm.MCacheSys)},
		{name: "MSpanInuse", value: float64(rtm.MSpanInuse)},
		{name: "MSpanSys", value: float64(rtm.MSpanSys)},
		{name: "Mallocs", value: float64(rtm.Mallocs)},
		{name: "NextGC", value: float64(rtm.NextGC)},
		{name: "NumForcedGC", value: float64(rtm.NumForcedGC)},
		{name: "NumGC", value: float64(rtm.NumGC)},
		{name: "OtherSys", value: float64(rtm.OtherSys)},
		{name: "PauseTotalNs", value: float64(rtm.PauseTotalNs)},
		{name: "StackInuse", value: float64(rtm.StackInuse)},
		{name: "StackSys", value: float64(rtm.StackSys)},
		{name: "Sys", value: float64(rtm.Sys)},
		{name: "TotalAlloc", value: float64(rtm.TotalAlloc)},
	}

	return Metrics
}
