package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func Alloc(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("Alloc", float64(memStats.Alloc))
}
