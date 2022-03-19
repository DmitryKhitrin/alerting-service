package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func LastGC(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("LastGC", float64(memStats.LastGC))
}
