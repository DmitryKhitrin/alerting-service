package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/common"
	"net/http"
)

type Service interface {
	StoreMetric(metric *common.Metrics) *common.Error
	GetMetric(metric *common.Metrics) (interface{}, *common.Error)
	GetTemplateWriter() (func(w http.ResponseWriter) error, error)
}
