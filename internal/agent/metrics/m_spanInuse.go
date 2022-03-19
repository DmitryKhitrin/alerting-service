package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func MSpanInuse(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("MSpanInuse", float64(memStats.MSpanInuse))
}
