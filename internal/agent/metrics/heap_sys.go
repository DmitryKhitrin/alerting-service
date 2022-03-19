package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func HeapSys(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("HeapSys", float64(memStats.HeapSys))
}
