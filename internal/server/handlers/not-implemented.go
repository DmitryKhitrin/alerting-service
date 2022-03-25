package handlers

import "net/http"

func NotImplemented(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "NotImplemented", http.StatusNotImplemented)
}
