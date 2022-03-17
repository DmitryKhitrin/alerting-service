package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func Frees(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("Frees", float64(memStats.Frees))
}
