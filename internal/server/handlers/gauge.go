package handlers

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/storage"
	"github.com/DmitryKhitrin/alerting-service/internal/server/utils"
	"log"
	"net/http"
	"strconv"
)

const GaugePath = "/update/gauge/"

func GaugeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	params := utils.ParseUrl(r.URL.Path)
	if value, err := strconv.ParseFloat(params.Value, 64); err == nil {
		storage.GetStorage().SetGauge(&storage.Gauge{
			Name:  params.Name,
			Value: value,
		})
	} else {
		log.Println(err)
	}

}
