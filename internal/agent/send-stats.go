package agent

import "github.com/DmitryKhitrin/alerting-service/internal/agent/metrics"

func SendStats() {

	storedMetrics := metrics.GetGaugeMetrics()
	for _, metric := range storedMetrics {
		request(metric.GetValue())
	}

	storedLocalMetrics := metrics.GetCounterMetrics()
	for _, localMetric := range storedLocalMetrics {
		request(localMetric.GetValue())
	}
}
