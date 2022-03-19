package metrics

import (
	"github.com/DmitryKhitrin/alerting-service/internal/utils"
)

func PullCount(pullCount int64) string {
	return utils.MakeStatStringCounter("PullCount", pullCount)
}
