package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func Sys(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("Sys", float64(memStats.Sys))
}
