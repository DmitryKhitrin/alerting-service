package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func TotalAlloc(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("TotalAlloc", float64(memStats.TotalAlloc))
}
