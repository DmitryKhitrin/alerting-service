package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func OtherSys(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("OtherSys", float64(memStats.OtherSys))
}
