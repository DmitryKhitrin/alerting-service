package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

func StackSys(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("StackSys", float64(memStats.StackSys))
}
