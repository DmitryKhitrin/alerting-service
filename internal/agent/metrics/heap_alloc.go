package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func HeapAlloc(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("HeapAlloc", float64(memStats.HeapAlloc))
}
