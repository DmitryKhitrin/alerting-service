package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func HeapReleased(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("HeapReleased", float64(memStats.HeapReleased))
}
