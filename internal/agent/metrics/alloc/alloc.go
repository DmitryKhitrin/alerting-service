package alloc

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
	"runtime"
)

const (
	Name = "Alloc"
)

func Get(memStats runtime.MemStats) string {
	return utils.MakeStatStringGauge(Name, float64(memStats.Alloc))
}
