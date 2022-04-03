package handlers

import "net/http"

func StatusNotImplemented(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "Method is not implemented yet", http.StatusNotImplemented)
}
