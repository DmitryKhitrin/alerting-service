package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func NumGC(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("NumGC", float64(memStats.NumGC))
}
