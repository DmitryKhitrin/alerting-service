package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func NextGC(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("NextGC", float64(memStats.NextGC))
}
