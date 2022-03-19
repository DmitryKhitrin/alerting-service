package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func StackInuse(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("StackInuse", float64(memStats.StackInuse))
}
