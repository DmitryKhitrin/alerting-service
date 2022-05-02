package agent

import (
	"bytes"
	"encoding/json"
	"github.com/DmitryKhitrin/alerting-service/internal/common"
	"log"
	"net/http"
)

const (
	protocol    = "http"
	contentType = "application/json"
)

type RequestService struct {
	Address string
}

func NewRequestService(address string) *RequestService {
	return &RequestService{Address: protocol + "://" + address}
}

func (r *RequestService) request(metric *common.Metrics) {
	jsonMetric, err := json.Marshal(metric)
	if err != nil {
		log.Println("error during marshaling in MetricSend %w", err)
		return
	}
	resp, err := http.Post(r.Address+"/update", contentType, bytes.NewBuffer(jsonMetric))

	if err != nil {
		log.Println(err)
		return
	}

	err = resp.Body.Close()
	if err != nil {
		log.Println(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Request error with %s, http status %d", metric.ID, resp.StatusCode)
	}
}
