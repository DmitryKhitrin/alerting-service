package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func PauseTotalNs(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("PauseTotalNs", float64(memStats.PauseTotalNs))
}
