package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func Mallocs(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("Mallocs", float64(memStats.Mallocs))
}
