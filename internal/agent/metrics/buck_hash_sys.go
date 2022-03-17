package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func BuckHashSys(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("BuckHashSys", float64(memStats.BuckHashSys))
}
