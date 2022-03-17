package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func GCSys(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("GCSys", float64(memStats.GCSys))
}
