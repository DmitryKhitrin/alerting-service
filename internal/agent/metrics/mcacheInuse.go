package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func MCacheInuse(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("MCacheInuse", float64(memStats.MCacheInuse))
}
