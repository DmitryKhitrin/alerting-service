package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func GCCPUFraction(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("GCCPUFraction", float64(memStats.GCCPUFraction))
}
