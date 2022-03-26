package agent

import (
	"log"
	"runtime"
)

func CollectStats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	log.Println("ticker CollectRuntimeMetrics")
}
