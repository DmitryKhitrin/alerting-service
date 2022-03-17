package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func HeapIdle(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("HeapIdle", float64(memStats.HeapIdle))
}
