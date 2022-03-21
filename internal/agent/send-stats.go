package agent

import "github.com/DmitryKhitrin/alerting-service/internal/agent/metrics"

func SendStats() {

	storedMetrics := metrics.CollectGauge()
	for _, metric := range storedMetrics {
		request(metric.GetValue())
	}

	storedLocalMetrics := metrics.CollectCounter()
	for _, localMetric := range storedLocalMetrics {
		request(localMetric.GetValue())
	}
}
