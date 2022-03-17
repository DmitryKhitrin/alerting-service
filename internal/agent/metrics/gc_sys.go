package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

//var a = []string{"GCSys"}

func GCSys(memStats *runtime.MemStats) string {
	return utils.MakeStatStringGauge("GCSys", float64(memStats.GCSys))
}
