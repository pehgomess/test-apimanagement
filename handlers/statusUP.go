package handlers

import "net/http"

func StatusUP(w http.ResponseWriter, r *http.Request) {
	//if r.URL.Path != "/health" {
	//	w.WriteHeader(http.StatusNotFound)
	//	w.Write([]byte("Asset not found\n"))
	//	return
	//}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API v1\n"))
}
