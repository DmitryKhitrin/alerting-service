package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func HeapInuse(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("HeapInuse", float64(memStats.HeapInuse))
}
