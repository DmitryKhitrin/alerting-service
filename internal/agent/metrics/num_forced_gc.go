package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func NumForcedGC(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("NumForcedGC", float64(memStats.NumForcedGC))
}
