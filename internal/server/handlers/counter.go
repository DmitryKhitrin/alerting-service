package handlers

import (
	"github.com/DmitryKhitrin/alerting-service/internal/server/storage"
	"github.com/DmitryKhitrin/alerting-service/internal/server/utils"
	"log"
	"net/http"
	"strconv"
)

const CounterPath = "/update/counter/"

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	params := utils.ParseUrl(r.URL.Path)
	if value, err := strconv.ParseInt(params.Value, 10, 64); err == nil {
		storage.GetStorage().SetCounter(&storage.Counter{
			Name:  params.Name,
			Value: value,
		})
		w.WriteHeader(http.StatusOK)
	} else {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
