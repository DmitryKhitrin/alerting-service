package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func MSpanSys(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("MSpanSys", float64(memStats.MSpanSys))
}
