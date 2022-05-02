package agent

import "github.com/DmitryKhitrin/alerting-service/internal/agent/metrics"

type StatsSender struct {
	Request *RequestService
}

func NewStatsSender(request *RequestService) *StatsSender {
	return &StatsSender{Request: request}
}

func (s *StatsSender) Send() {

	storedMetrics := metrics.GetGaugeMetrics()
	for _, metric := range storedMetrics {
		s.Request.request(metric.GetValue())
	}

	storedLocalMetrics := metrics.GetCounterMetrics()
	for _, localMetric := range storedLocalMetrics {
		s.Request.request(localMetric.GetValue())
	}
}
