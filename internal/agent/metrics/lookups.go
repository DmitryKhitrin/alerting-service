package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func Lookups(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("Lookups", float64(memStats.Lookups))
}
