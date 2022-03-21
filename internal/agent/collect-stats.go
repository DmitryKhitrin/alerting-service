package agent

import "runtime"

func CollectStats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
}
