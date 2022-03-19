package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func HeapObjects(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("HeapObjects", float64(memStats.HeapObjects))
}
