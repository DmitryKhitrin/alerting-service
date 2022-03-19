package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
)

func RandomValue(randomValue int64) string {
	return utils.MakeStatStringCounter("RandomValue", randomValue)
}
