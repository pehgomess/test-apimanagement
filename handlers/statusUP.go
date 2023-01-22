package handlers

import "net/http"

func StatusUP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/status" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Asset not found\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1\n"))
}
