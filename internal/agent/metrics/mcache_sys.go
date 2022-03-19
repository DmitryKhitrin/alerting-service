package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func MCacheSys(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("MCacheSys", float64(memStats.LastGC))
}
